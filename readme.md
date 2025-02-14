to build and push image to minikube

```sh
eval $(minikube docker-env)
docker build -f Dockerfile-client -t grpc-client:latest .
docker build -f Dockerfile-server -t grpc-server:latest .

minikube kubectl -- apply -f grpc-client-deployment.yaml
minikube kubectl -- apply -f grpc-client-service.yaml
minikube kubectl -- apply -f grpc-server-deployment.yaml
minikube kubectl -- apply -f grpc-server-service.yaml

# install istioctl
curl -sL https://istio.io/downloadIstioctl | sh -
export PATH=$HOME/.istioctl/bin:$PATH

kubectl config use-context minikube
istioctl install

minikube kubectl -- label namespace default istio-injection=enabled --overwrite

$ kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.24/samples/addons/prometheus.yaml
istioctl dashboard prometheus
```
