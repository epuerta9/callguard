import logging

# Setup logging
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)


def get_list_contacts(account_id: str, customer_phone: str) -> str:
    """
    Simulates fetching a contact list for a given account and customer phone number.
    This is a mock implementation using hardcoded data.
    """

    logger.info(
        f"Fetching contacts for account: {account_id} and phone: {customer_phone}")

    # Example hardcoded contacts
    contacts = [
        {
            "name": "John Doe",
            "phone": "+17138342650",
            "email": "john.doe@example.com"
        },
        {
            "name": "Maria Lopez",
            "phone": "+18185551234",
            "email": "maria.lopez@example.com"
        },
        {
            "name": "Carlos Martinez",
            "phone": "+14155559876",
            "email": "carlos.martinez@example.com"
        }
    ]

    # Generate the response text
    response_lines = [f"{idx+1}. {contact['name']} - {contact['phone']} - {contact['email']}"
                      for idx, contact in enumerate(contacts)]

    response_text = f"The contacts for account {account_id} are:\n" + "\n".join(
        response_lines)

    logger.info(f"Contact list generated: {response_text}")

    return response_text
