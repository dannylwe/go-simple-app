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

### creating new helm char
- helm create <chart-name>