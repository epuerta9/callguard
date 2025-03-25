from typing import List, Optional
from pydantic import BaseModel


class Voice(BaseModel):
    voice_id: str
    provider: str = "11labs"
    stability: float = 0.5
    similarity_boost: float = 0.75
    model: str = "eleven_multilingual_v2"
    filler_injection_enabled: bool = False


class Message(BaseModel):
    role: str
    content: str


class Model(BaseModel):
    model: str = "gpt-4o-mini"
    tool_ids: List[str]
    messages: List[Message]
    provider: str = "openai"
    emotion_recognition_enabled: bool = False


class Transcriber(BaseModel):
    model: str = "nova-2"
    language: str = "multi"


class Assistant(BaseModel):
    voice: Voice
    model: Model
    recording_enabled: bool = True
    first_message: str
    voicemail_message: str
    end_call_function_enabled: bool = True
    end_call_message: str
    transcriber: Transcriber


class AssistantResponse(BaseModel):
    assistant: Assistant


class CustomerInfo(BaseModel):
    number: str


class CallInfo(BaseModel):
    customer: CustomerInfo


class PhoneNumber(BaseModel):
    number: str


class MessageRequest(BaseModel):
    call: CallInfo
    phone_number: PhoneNumber


class AssistantRequest(BaseModel):
    message: MessageRequest
