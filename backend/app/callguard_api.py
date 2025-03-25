import json
from typing import List, Optional
from pydantic import BaseModel
import httpx
import os
from dotenv import load_dotenv

load_dotenv()


class Assistant(BaseModel):
    id: str
    name: str


class AssistantData(BaseModel):
    account_id: str
    company_name: str
    credits: float
    lang: str
    tools: str
    assistants: str
    system_prompt: str


class CallGuardAPIError(Exception):
    pass


class NoCreditsError(CallGuardAPIError):
    pass


class SpamError(CallGuardAPIError):
    pass


class CallGuardAPI:
    def __init__(self):
        self.base_url = os.getenv(
            "CALLGUARD_API_URL", "https://api.callguard.ai")
        self.api_key = os.getenv("CALLGUARD_API_KEY") or "XYZ"
        if not self.api_key:
            raise ValueError(
                "CALLGUARD_API_KEY environment variable is not set")

    async def get_assistant_data(self, phone_number: str) -> AssistantData:
        """Get assistant data for a given phone number."""
        async with httpx.AsyncClient() as client:
            response = await client.get(
                f"{self.base_url}/assistant_data",
                params={"phone_number": phone_number},
                headers={"Authorization": f"Bearer {self.api_key}"}
            )

            if response.status_code != 200:
                raise CallGuardAPIError(
                    f"Error getting assistant data: {response.text}")

            return AssistantData(**response.json())

    async def check_spam(self, phone_number: str) -> bool:
        """Check if a phone number is marked as spam."""
        async with httpx.AsyncClient() as client:
            response = await client.get(
                f"{self.base_url}/check_spam",
                params={"phone_number": phone_number},
                headers={"Authorization": f"Bearer {self.api_key}"}
            )

            if response.status_code != 200:
                raise CallGuardAPIError(
                    f"Error checking spam: {response.text}")

            return response.json().get("is_spam", False)
