import torch
from fastapi import FastAPI
from transformers import AutoTokenizer, AutoModelForCausalLM

from settings import MODEL_NAME, DEVICE
from utils import get_model_params, create_prompt, tokenize_prompt

app = FastAPI()

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME, device_map="auto", cache_dir=".cache")


@app.post("/runsync")
async def run_model(job: dict):
    prompt = create_prompt(job["input"]["prompts"])
    inputs = tokenize_prompt(tokenizer, prompt, DEVICE)
    max_new_tokens, repetition_penalty = get_model_params(job)

    with torch.no_grad():
        outputs = model.generate(
            inputs, max_new_tokens=max_new_tokens, repetition_penalty=repetition_penalty
        )
    output = outputs[0][len(inputs[0]):]
    model_output_decoded = tokenizer.decode(output, skip_special_tokens=True)
    return {"output": model_output_decoded}


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=8000)
