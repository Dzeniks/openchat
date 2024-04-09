# Builder stage
FROM python:3.11.7-slim-bookworm AS builder

WORKDIR /app
COPY run-cpu-local.py run-cpu-local.py
COPY utils.py utils.py
COPY settings.py settings.py
COPY requirements.txt requirements.txt

ENV MODEL_NAME=google/gemma-2b-it

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["python",  "run-cpu-local.py"]