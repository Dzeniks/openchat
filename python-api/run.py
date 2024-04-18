
import runpod
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, BitsAndBytesConfig
from time import perf_counter_ns

from utils import get_model_params, create_prompt, tokenize_prompt
from settings import DEVICE, MODEL_NAME

nf4_config = BitsAndBytesConfig(
    load_in_4bit=True
)

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME,
                                             torch_dtype=torch.bfloat16,
                                             quantization_config=nf4_config,
                                             attn_implementation="flash_attention_2",
                                             device_map="auto", cache_dir=".cache")


def run_model(job):
    start = perf_counter_ns()
    prompt = create_prompt(job["input"]["prompts"])
    inputs = tokenize_prompt(tokenizer, prompt, DEVICE)
    max_new_tokens, repetition_penalty = get_model_params(job)
    with torch.no_grad():
        outputs = model.generate(
            input_ids=inputs, max_new_tokens=max_new_tokens, repetition_penalty=repetition_penalty
        )
    output = outputs[0][len(inputs[0]):]
    model_output_decoded = tokenizer.decode(output, skip_special_tokens=True)
    # time
    end = perf_counter_ns()
    print(f"Time taken: {(end - start) / 1e6:.2f} ms")
    return model_output_decoded


runpod.serverless.start({"handler": run_model})
