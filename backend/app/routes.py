from fastapi import APIRouter, Request
from app.services import assistant_request_handler

router = APIRouter()


@router.get("/")
async def home():
    return {"success": True, "message": "CallGuard API v1.0.0"}


@router.post("/assistant_request")
async def assistant_request(req: Request):
    body = await req.body()
    return await assistant_request_handler(body)
    return Response(status_code=200)
