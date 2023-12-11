# count-client

## Usage
Ensure the count-server is running, then run the container:
```
❯ docker build -t count-client .
❯ docker run --network=count-network -d count-client
```
