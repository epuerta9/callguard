# CallGuard VAPI Assistant Service

A Python-based service that handles VAPI (Voice API) requests for the CallGuard platform. This service manages assistant interactions, spam detection, and credit management for voice calls.

## Features

- Assistant request handling
- Spam detection
- Credit management
- Multi-language support (English/Spanish)
- Development mode for testing without API credentials

## Prerequisites

- Python 3.8+
- Virtual environment (venv)
- CallGuard API credentials (for production use)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/callguard.git
cd callguard
```

2. Create and activate a virtual environment:
```bash
python -m venv venv
source venv/bin/activate  # On Windows: .\venv\Scripts\activate
```

3. Install dependencies:
```bash
pip install -r requirements.txt
```

## Configuration

For production use, set up the following environment variables:
```bash
CALLGUARD_API_KEY=your_api_key_here
CALLGUARD_API_URL=https://api.callguard.ai  # Replace with actual API URL
```

For development, the service will run in development mode with mock data if these variables are not set.

## Running the Service

Start the service using uvicorn:
```bash
uvicorn main:app --reload --port 8001
```

The service will be available at `http://127.0.0.1:8001`.

## API Endpoints

### GET /
Health check endpoint that returns service version information.

### POST /assistant_request
Handles assistant requests for voice calls.

Example request:
```json
{
  "message": {
    "call": {
      "customer": {
        "number": "+1234567890"
      }
    },
    "phone_number": {
      "number": "+1987654321"
    }
  }
}
```

Example response:
```json
{
  "assistant": {
    "voice": {
      "voice_id": "mock_voice_id",
      "provider": "11labs",
      "stability": 0.5,
      "similarity_boost": 0.75,
      "model": "eleven_multilingual_v2",
      "filler_injection_enabled": false
    },
    "model": {
      "model": "gpt-4o-mini",
      "tool_ids": ["get_list_contacts", "check_slots", "confirm_appointment"],
      "messages": [
        {
          "role": "system",
          "content": "..."
        }
      ],
      "provider": "openai",
      "emotion_recognition_enabled": false
    },
    "recording_enabled": true,
    "first_message": "Hello, this is Alex from CallGuard Demo, how can I help you today?",
    "voicemail_message": "You've reached Alex's voicemail...",
    "end_call_function_enabled": true,
    "end_call_message": "Thank you for calling CallGuard Demo, have a great day! Goodbye",
    "transcriber": {
      "model": "nova-2",
      "language": "multi"
    }
  }
}
```

## Development

### Project Structure
```
callguard/
├── app/
│   ├── __init__.py
│   ├── models.py       # Pydantic models for request/response
│   ├── routes.py       # FastAPI route definitions
│   ├── services.py     # Business logic
│   └── callguard_api.py # API client for CallGuard
├── main.py            # Application entry point
├── requirements.txt   # Project dependencies
└── README.md         # This file
```

### Adding New Features

1. Define new models in `app/models.py`
2. Implement business logic in `app/services.py`
3. Add new routes in `app/routes.py`
4. Update API client in `app/callguard_api.py` if needed

### Testing

You can test the service using curl:
```bash
curl -X POST http://127.0.0.1:8001/assistant_request \
-H "Content-Type: application/json" \
-d '{
  "message": {
    "call": {
      "customer": {
        "number": "+1234567890"
      }
    },
    "phone_number": {
      "number": "+1987654321"
    }
  }
}'
```

## License

[Add your license information here]
