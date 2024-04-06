import os
import dotenv

dotenv.load_dotenv()

model_name = os.environ.get("MODEL_NAME")

if model_name is None:
    raise Exception("MODEL_NAME environment variable is not set")
print(model_name)
