from typing import List, Dict, Any

from settings import MAX_NEW_TOKENS, REPEAT_PENALTY


def create_prompt(prompts: list[dict]) -> list[dict[str, str]]:
    chat = []
    for prompt in prompts:
        if prompt.get("SenderID", None) == "AI":
            chat.append({"role": "model", "content": prompt["Content"]})
        else:
            chat.append({"role": "user", "content": prompt["Content"]})
    return chat


def tokenize_prompt(tokenizer, chat: list[dict[str, str]], device="cpu"):
    formated_chat = tokenizer.apply_chat_template(chat, tokenize=False, add_generation_prompt=True)
    print(formated_chat)
    return tokenizer.encode(formated_chat, return_tensors="pt").to(device)


def get_model_params(job):
    max_new_tokens = job["input"].get("max_new_tokens", MAX_NEW_TOKENS)
    repetition_penalty = job["input"].get("repetition_penalty", REPEAT_PENALTY)
    return max_new_tokens, repetition_penalty
