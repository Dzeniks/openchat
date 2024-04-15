import torch
from peft import get_peft_model, LoraConfig, TaskType
from transformers import AutoModelForCausalLM

from settings import BASE_MODEL

LoRaConfig = LoraConfig(
    r=64,
    lora_alpha=16,
    lora_dropout=0.1,
    task_type=TaskType.CAUSAL_LM,
)

def get_model_peft():
    base_model = AutoModelForCausalLM.from_pretrained(BASE_MODEL, torch_dtype=torch.float16)
    return get_peft_model(base_model, LoRaConfig)
