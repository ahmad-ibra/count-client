# count-client
- [Companion server](https://github.com/ahmad-ibra/count-server)
- [Running on k8s](https://github.com/ahmad-ibra/count-k8s)

## Usage
Ensure the count-server is running, then run the container:
```
❯ docker build -t count-client .
❯ docker run --network=count-network -d count-client
```
