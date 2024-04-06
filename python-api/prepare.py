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

tokenizer = AutoTokenizer.from_pretrained(model_name, cache_dir=".cache")
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.float16,
                                             device_map="auto", cache_dir=".cache")