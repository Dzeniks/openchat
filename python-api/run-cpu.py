import os
import dotenv
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
import runpod

from utils import get_model_params, create_prompt, tokenize_prompt

dotenv.load_dotenv()

DEVICE = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model_name = os.getenv("MODEL_NAME")

if model_name is None:
    raise Exception("MODEL_NAME environment variable is not set")

print(f"Using device: {DEVICE} and model: {model_name}")
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.float16,
                                             device_map="auto")


def run_model(job):
    print(f"Using device: {DEVICE} and model: {model_name}")
    prompt = create_prompt(job["input"]["prompt"])
    inputs = tokenize_prompt(tokenizer, prompt, DEVICE)
    model_params = get_model_params(job)
    with torch.no_grad():
        outputs = model.generate(
            **inputs, **model_params
        )
    return tokenizer.decode(outputs[0], skip_special_tokens=True)


runpod.serverless.start({"handler": run_model})
