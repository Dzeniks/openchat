# openchat

For local development, you can use the following command to start the server:

I recommend using a virtual environment to run the server:
For linux
```bash
python3 -m venv venv
source venv/bin/activate
```
For windows
```bash
python -m venv venv
venv\Scripts\activate
```

then you can install the requirements and start the server:

```bash
cd ./python-api
pip install -r requirements.txt
python run.py --rp_server_api
```

then you can run docker-compose up to start the server and the client:

```bash
docker-compose build
docker-compose up
```
