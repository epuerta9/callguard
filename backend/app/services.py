import json
import random
import os
from datetime import datetime
import pytz
from fastapi.responses import JSONResponse
from app.models import AssistantRequest, AssistantResponse, Voice, Model, Message, Assistant, Transcriber
from app.callguard_api import CallGuardAPI, SpamError, NoCreditsError

# Initialize API client only if credentials are available
try:
    api_client = CallGuardAPI()
except ValueError:
    print("Warning: CALLGUARD_API_KEY not set, running in development mode")
    api_client = None

async def assistant_request_handler(body: bytes) -> JSONResponse:
    try:
        # Parse request
        request_data = json.loads(body)
        assistant_request = AssistantRequest(**request_data)
        
        if api_client is not None:
            # Check for spam
            if await api_client.check_spam(assistant_request.message.call.customer.number):
                return JSONResponse(
                    status_code=400,
                    content={"error": "Number marked as spam"}
                )
            
            # Get assistant data
            assistant_data = await api_client.get_assistant_data(assistant_request.message.phone_number.number)
            
            # Check credits
            if assistant_data.credits < 0.30:
                return JSONResponse(
                    status_code=400,
                    content={"error": "Insufficient credits"}
                )
            
            # Parse assistants
            assistants = json.loads(assistant_data.assistants)
            selected_assistant = random.choice(assistants)
            tools_list = assistant_data.tools.strip('{}').split(',')
            company_name = assistant_data.company_name
            lang = assistant_data.lang
            account_id = assistant_data.account_id
            system_prompt = assistant_data.system_prompt
        else:
            # Development mode - use mock data
            selected_assistant = {
                "id": "mock_voice_id",
                "name": "Alex"
            }
            tools_list = ["get_list_contacts", "check_slots", "confirm_appointment"]
            company_name = "CallGuard Demo"
            lang = "english"
            account_id = "demo_account"
            system_prompt = "You are a helpful AI assistant."
        
        # Get current time in Chicago timezone
        chicago_tz = pytz.timezone('America/Chicago')
        current_time = datetime.now(chicago_tz)
        
        # Prepare messages based on language
        first_message = (
            f"Hola, soy {selected_assistant['name']} de {company_name}, en que puedo ayudarte hoy?"
            if lang == "spanish"
            else f"Hello, this is {selected_assistant['name']} from {company_name}, how can I help you today?"
        )
        
        end_call_message = (
            f"Gracias por llamar a {company_name}, que tengas un buen día! Adios"
            if lang == "spanish"
            else f"Thank you for calling {company_name}, have a great day! Goodbye"
        )
        
        voicemail_message = (
            f"Has llegado al buzón de voz de {selected_assistant['name']}. Por favor deja un mensaje después del pitido, y te responderemos lo antes posible."
            if lang == "spanish"
            else f"You've reached {selected_assistant['name']}'s voicemail. Please leave a message after the beep, and we'll get back to you as soon as possible."
        )
        
        # Create response
        response = AssistantResponse(
            assistant=Assistant(
                voice=Voice(voice_id=selected_assistant['id']),
                model=Model(
                    tool_ids=tools_list,
                    messages=[
                        Message(
                            role="system",
                            content=(
                                f"You named is {selected_assistant['name']}\n"
                                f"You can speak English and Spanish; the default language is {lang}, use the language that the user speak all time.\n"
                                f"the AccountID is {account_id}\n"
                                f"the current time is {current_time.strftime('%A, %d %B %Y %H:%M:%S %z')}\n"
                                "always say goodbye\n"
                                f"{system_prompt}"
                            )
                        )
                    ]
                ),
                first_message=first_message,
                voicemail_message=voicemail_message,
                end_call_message=end_call_message,
                transcriber=Transcriber()
            )
        )
        
        return JSONResponse(response.model_dump())
        
    except SpamError:
        return JSONResponse(
            status_code=400,
            content={"error": "Number marked as spam"}
        )
    except NoCreditsError:
        return JSONResponse(
            status_code=400,
            content={"error": "Insufficient credits"}
        )
    except Exception as e:
        return JSONResponse(
            status_code=500,
            content={"error": str(e)}
        )
