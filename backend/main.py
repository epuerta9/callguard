import logging
from fastapi import FastAPI, Request, Response
from mangum import Mangum
from app.routes import router as app_router

logging.basicConfig(level=logging.INFO)

app = FastAPI(title="CallGuard v1.0.0")
app.include_router(app_router)

handler = Mangum(app)
