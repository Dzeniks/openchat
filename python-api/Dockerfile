# Builder stage
FROM python:3.11.7-bookworm AS builder

WORKDIR /app
COPY . .
COPY settings.py .env


# Install dependencies and prepare environment
RUN pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt

EXPOSE 8000
CMD ["python", "-u", "run.py"]
