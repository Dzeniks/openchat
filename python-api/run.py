
import runpod
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer, BitsAndBytesConfig

from utils import get_model_params, create_prompt, tokenize_prompt
from settings import DEVICE, MODEL_NAME

nf4_config = BitsAndBytesConfig(
    load_in_4bit=True,
    bnb_4bit_compute_dtype=torch.float16,
    bnb_4bit_use_double_quant=True,
    bnb_4bit_quant_type="nf4"
)

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME,
                                             torch_dtype=torch.float16,
                                             quantization_config=nf4_config,
                                             device_map="auto", cache_dir=".cache")


def run_model(job):
    with torch.backends.cuda.sdp_kernel(
            enable_flash=True,
            enable_math=False,
            enable_mem_efficient=False
    ):
        prompt = create_prompt(job["input"]["prompts"])
        print(prompt)
        inputs = tokenize_prompt(tokenizer, prompt, DEVICE)
        max_new_tokens, repetition_penalty = get_model_params(job)
        with torch.no_grad():
            outputs = model.generate(
                inputs, max_new_tokens=max_new_tokens, repetition_penalty=repetition_penalty
            )
        output = outputs[0][len(inputs[0]):]
        model_output_decoded = tokenizer.decode(output, skip_special_tokens=True)
    return model_output_decoded


runpod.serverless.start({"handler": run_model})
