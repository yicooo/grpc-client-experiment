to build and push image to minikube

```sh
eval $(minikube docker-env)
docker build -f Dockerfile-client -t grpc-client:latest .
docker build -f Dockerfile-server -t grpc-server:latest .

minikube kubectl -- apply -f grpc-service.yaml
minikube kubectl -- apply -f grpc-server-deployment.yaml
minikube kubectl -- apply -f grpc-client.yaml
```
