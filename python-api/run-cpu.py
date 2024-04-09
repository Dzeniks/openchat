from settings import DEVICE, MODEL_NAME

import runpod
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

from utils import get_model_params, create_prompt, tokenize_prompt


if MODEL_NAME is None:
    raise Exception("MODEL_NAME environment variable is not set")

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME, torch_dtype=torch.float16,
                                             device_map="auto", cache_dir=".cache")


def run_model(job: dict):
    prompt = create_prompt(job["input"]["prompts"])
    inputs = tokenize_prompt(tokenizer, prompt, DEVICE)
    max_new_tokens, repetition_penalty = get_model_params(job)
    with torch.no_grad():
        outputs = model.generate(
            inputs, max_new_tokens=max_new_tokens, repetition_penalty=repetition_penalty
        )
    output = outputs[0][len(inputs[0]):]
    print(tokenizer.decode(outputs[0], skip_special_tokens=True))
    model_output_decoded = tokenizer.decode(output, skip_special_tokens=True)
    return model_output_decoded


runpod.serverless.start({"handler": run_model})
