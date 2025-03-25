def make_response(tool_call_id: str, result: str):
    return {
        "results": [
            {
                "toolCallId": tool_call_id,
                "result": result
            }
        ]
    }
