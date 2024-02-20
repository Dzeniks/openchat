import os
import dotenv
import torch
from transformers import AutoModel, AutoTokenizer, BitsAndBytesConfig
import runpod

dotenv.load_dotenv()


DEVICE = torch.device("cuda" if torch.cuda.is_available() else "cpu")
model_name = os.getenv("MODEL_NAME")

if model_name is None:
    raise Exception("MODEL_NAME environment variable is not set")

model = AutoModel.from_pretrained(model_name)

