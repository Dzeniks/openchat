from settings import PROMPT_TEMPLATE_LAMMA, MESSAGE_SYMBOL, MAX_NEW_TOKENS, REPEAT_PENALTY


def create_prompt(prompts: list[dict]) -> str:
    full_prompt = ""
    for prompt in prompts:
        if prompt.get("SenderID", None) == "AI":
            full_prompt += prompt["Content"] + "</s>"
        else:
            full_prompt += PROMPT_TEMPLATE_LAMMA.replace(MESSAGE_SYMBOL, prompt["Content"])
    return full_prompt


def tokenize_prompt(tokenizer, prompt: str, device="cpu"):
    return tokenizer.encode(prompt, return_tensors="pt").to(device)


def get_model_params(job):
    max_new_tokens = job["input"].get("max_new_tokens", MAX_NEW_TOKENS)
    repetition_penalty = job["input"].get("repetition_penalty", REPEAT_PENALTY)
    return max_new_tokens, repetition_penalty