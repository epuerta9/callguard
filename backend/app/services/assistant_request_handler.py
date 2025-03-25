import json
import requests
import logging

logger = logging.getLogger(__name__)

VAPI_API_KEY = "83378317-e7a4-4fb4-b429-cfd9af3104db"
VAPI_BASE_URL = "https://api.vapi.ai"

async def assistant_request_handler(body: bytes):
    """
    Handle assistant request and create an assistant using VAPI API
    """
    try:
        # Parse the request body
        data = json.loads(body)
        message = data.get('message', {})
        
        # Check if it's an assistant request
        if message.get('type') != 'assistant-request':
            logger.warning(f"Received non-assistant request type: {message.get('type')}")
            return {"error": "Invalid request type"}
        
        # Extract call and phone number information
        call = message.get('call', {})
        phone_number = message.get('phoneNumber', {})
        
        # Prepare the assistant configuration
        assistant_config = {
            "name": "CallGuard Assistant",
            "description": "A helpful assistant for handling phone calls",
            "model": "gpt-4-turbo-preview",
            "system_prompt": """You are a helpful assistant that handles phone calls professionally.
            You should be courteous, clear, and efficient in your communication.
            Always identify yourself as CallGuard Assistant and ask how you can help.""",
            "voice_config": {
                "provider": "elevenlabs",
                "voice_id": "default",
                "stability": 0.5,
                "similarity_boost": 0.75
            }
        }
        
        # Create the assistant using the VAPI API
        response = requests.post(
            f"{VAPI_BASE_URL}/assistant",
            headers={
                "Authorization": f"Bearer {VAPI_API_KEY}",
                "Content-Type": "application/json"
            },
            json=assistant_config
        )
        
        # Check if the request was successful
        response.raise_for_status()
        
        # Return the assistant response
        return {
            "success": True,
            "assistant": response.json()
        }
        
    except json.JSONDecodeError as e:
        logger.error(f"Failed to parse request body: {e}")
        return {"error": "Invalid JSON in request body"}
    except requests.exceptions.RequestException as e:
        logger.error(f"Error creating assistant: {e}")
        return {"error": f"Failed to create assistant: {str(e)}"}
    except Exception as e:
        logger.error(f"Error handling assistant request: {e}")
        return {"error": str(e)} 