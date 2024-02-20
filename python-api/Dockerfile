# Build stage
FROM python:3.11.7-bookworm

WORKDIR /app
COPY . .

RUN pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt \
    && pip install python-dotenv \
    && python prepare.py

EXPOSE 8000
CMD ["python", "-u", "/run.py"]
#CMD ["python", "run.py", "--rp_serve_api"]

