from fastapi import APIRouter, Request
import json
import os
from datetime import datetime
from app.services import assistant_request_handler

router = APIRouter()


def save_request_body(body, prefix="request"):
    """Save the request body as a JSON file with timestamp."""
    # Create logs directory if it doesn't exist
    logs_dir = os.path.join(os.path.dirname(os.path.dirname(__file__)), "logs")
    os.makedirs(logs_dir, exist_ok=True)
    
    # Generate a timestamp for the filename
    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    log_path = os.path.join(logs_dir, f"{prefix}_{timestamp}.json")
    
    # Save the request body as JSON
    try:
        # Try to parse as JSON first
        json_body = json.loads(body)
        with open(log_path, "w") as f:
            json.dump(json_body, f, indent=2)
    except json.JSONDecodeError:
        # If not valid JSON, save as string in a JSON file
        with open(log_path, "w") as f:
            json.dump({"raw_body": body.decode("utf-8", errors="replace")}, f, indent=2)
    
    print(f"[DEBUG] Saved {prefix} body to: {log_path}")
    return log_path


@router.get("/")
async def home():
    return {"success": True, "message": "CallGuard API v1.0.0"}


@router.post("/")
async def home(req: Request):
    body = await req.body()
    print(f"[DEBUG] Received POST request body: {body}")
    log_path = save_request_body(body)
    
    # Parse the request body to check the message type
    try:
        data = json.loads(body)
        message = data.get('message', {})
        
        # If it's an assistant request, handle it
        if message.get('type') == 'assistant-request':
            return await assistant_request_handler(body)
            
    except json.JSONDecodeError:
        return {"error": "Invalid JSON in request body"}
    
    return {"success": True, "message": "CallGuard API v1.0.0"}


@router.post("/assistant_request")
async def assistant_request(req: Request):
    body = await req.body()
    print(f"[DEBUG] Received assistant request body: {body}")
    log_path = save_request_body(body, prefix="assistant_request")
    return await assistant_request_handler(body)
