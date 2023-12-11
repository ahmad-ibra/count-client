# count-client
- [Companion server](https://github.com/ahmad-ibra/count-server)
- [Running on k8s](https://github.com/ahmad-ibra/count-k8s)

## Usage
Ensure the count-server is running, then run the container:
```
❯ docker build -t count-client .
❯ docker run -e HOST=count-server -e PORT=1234 --network=count-network -d count-client
```

## Useful Commands
Push new version of the image:
```
❯ docker build -t ahmadibraspectrocloud/count-client:latest .
❯ docker push ahmadibraspectrocloud/count-client:latest
```
