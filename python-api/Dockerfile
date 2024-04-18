FROM pytorch/pytorch:2.2.2-cuda12.1-cudnn8-devel

RUN apt-get update && apt-get install -y \
    git

WORKDIR /app
COPY requirements.txt requirements.txt

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# Install latest cuda and set CUDA_HOME


COPY run.py run.py
COPY utils.py utils.py
COPY settings.py settings.py
COPY .cache .cache
COPY .env .env
COPY test_input.json test_input.json

EXPOSE 8000

CMD ["python", "-u", "run.py"]
