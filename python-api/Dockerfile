# Build stage
FROM python:3.11.7-bookworm as builder
WORKDIR /
COPY requirements.txt /
RUN pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt
# Add prepare.py
ENV MODEL_NAME=albert/albert-base-v2
COPY prepare.py /
RUN python prepare.py

# Final stage
FROM python:3.11.7-bookworm
WORKDIR /
COPY --from=builder / /
COPY run.py /
CMD ["python", "-u", "/run.py"]
