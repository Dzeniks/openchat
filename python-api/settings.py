from dotenv import load_dotenv

load_dotenv()

DEVICE = "cuda" if torch.cuda.is_available() else "cpu"
MODEL_NAME = os.getenv("MODEL_NAME")

if MODEL_NAME is None:
    raise Exception("MODEL_NAME environment variable is not set")


MESSAGE_SYMBOL = "message"
PROMPT_TEMPLATE_LAMMA = f"<s>[INST]{MESSAGE_SYMBOL} [/INST]"

MAX_NEW_TOKENS = 100
REPEAT_PENALTY = 1.0
