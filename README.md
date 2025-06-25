### start application with Docker
- docker compose build
- docker compose up -d

### start application binary
- go build -o app
- ./app

### run tests
- go run tests -cover

### run k8s
- alias k=kubectl
- minikube ip (put this in ingress.yaml)
- k apply -f deployment.yaml
- k apply -f service.yaml
- k apply -f ingress.yaml

### building docker image with multiple tags
- docker build -t acehanks/go-simple-app:v2 -t acehanks/go-simple-app:latest .
- docker push acehanks/go-simple-app --all-tags

### creating new helm char
- helm create <chart-name>

### create deployment with helm
- helm install go-release simple-go-app-helm

### how to upgrade a deployment
helm upgrade go-release simple-go-app-helm --values=simple-go-app-helm/values.yaml --atomic

## use templating override
- kubectl create namespace dev
- helm install go-release-dev simple-go-app-helm --values=simple-go-app-helm/values.yaml -f simple-go-app-helm/values.dev.yaml


### NOTE
Whatever is put in values.*.yaml will override the contents of values.yaml