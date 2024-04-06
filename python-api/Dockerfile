# Builder stage
FROM python:3.11.7-slim-bookworm AS builder

WORKDIR /app
COPY run.py run.py
COPY utils.py utils.py
COPY settings.py settings.py
COPY requirements.txt requirements.txt
COPY .cache .cache

COPY test_input.json test_input.json
COPY .env .env


RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["python", "-u", "run.py"]
