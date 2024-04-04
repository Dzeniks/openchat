# Builder stage
FROM python:3.11.7-bookworm AS builder

WORKDIR /app
COPY . .
COPY .env .env
COPY settings.py settings.py
COPY requirements.txt requirements.txt

# Install dependencies and prepare environment
RUN python -m venv venv
RUN /bin/bash -c "source venv/bin/activate \
    && pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt \
    && python prepare.py"

# Final image
FROM python:3.11.7-slim-bookworm

WORKDIR /app

COPY --from=builder /app /app
COPY --from=builder /root/.cache/huggingface/hub /root/.cache/huggingface/hub
COPY --from=builder /app/.env .env
COPY --from=builder /app/settings.py settings.py

EXPOSE 8000

CMD ["./venv/bin/python", "-u", "run.py"]
