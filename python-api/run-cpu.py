import os
import dotenv
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, BitsAndBytesConfig
import runpod

# Import GemmaTokenizer and GemmaModel


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
    prompt = job["input"]["prompt"]
    max_new_tokens = job["input"].get("max_new_tokens", 100)
    repetition_penalty = job["input"].get("repetition_penalty", 1.0)
    inputs = tokenizer(prompt, return_tensors="pt").to(DEVICE)

    with torch.no_grad():
        outputs = model.generate(
            **inputs, max_new_tokens=max_new_tokens, repetition_penalty=repetition_penalty
        )
        
    return tokenizer.decode(outputs[0], skip_special_tokens=True)

runpod.serverless.start({"handler": run_model})
