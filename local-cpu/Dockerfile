# Builder stage
FROM python:3.11.7-slim-bookworm AS builder

WORKDIR /app
COPY . .

RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 8000

CMD ["python",  "run-cpu-local.py"]