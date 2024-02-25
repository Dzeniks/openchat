# Builder stage
FROM python:3.11.7-bookworm AS builder

WORKDIR /app
COPY . .
COPY .env .env


# Install dependencies and prepare environment
RUN pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt \
    && python prepare.py

# Final stage
FROM python:3.11.7-slim
#
WORKDIR /app

# Copy python dependencies from the builder stage
COPY --from=builder /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
# Copy built artifacts from the builder stage
COPY --from=builder /app .
COPY --from=builder /root/.cache/huggingface/ /root/.cache/huggingface/
# Copy transformers .cache from the builder stage

EXPOSE 8000
CMD ["python", "-u", "run.py"]
