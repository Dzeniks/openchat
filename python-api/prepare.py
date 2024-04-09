import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
import runpod

from utils import get_model_params, create_prompt, tokenize_prompt
from settings import MODEL_NAME

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME, torch_dtype=torch.float16,
                                             device_map="auto", cache_dir=".cache")