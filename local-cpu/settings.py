import os

import torch
from dotenv import load_dotenv

load_dotenv()

DEVICE = "cuda" if torch.cuda.is_available() else "cpu"
MODEL_NAME = os.getenv("MODEL_NAME")

if MODEL_NAME is None:
    raise Exception("MODEL_NAME environment variable is not set")

MAX_NEW_TOKENS = 4000
REPEAT_PENALTY = 1.0
