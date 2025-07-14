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

### creating new helm chart
- helm create <chart-name>

### create deployment with helm
- helm install go-release simple-go-app-helm

### how to upgrade a deployment
helm upgrade go-release simple-go-app-helm --values=simple-go-app-helm/values.yaml --atomic

## use templating override
- kubectl create namespace dev
- helm install go-release-dev simple-go-app-helm --values=simple-go-app-helm/values.yaml -f simple-go-app-helm/values.dev.yaml

## How to rollback on helm
- helm rollback go-simple-release 3

### NOTE
Whatever is put in values.*.yaml will override the contents of values.yaml

## Install argoCD
- k create namespace argocd
- k apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/refs/heads/master/manifests/install.yaml
- port forward argocd -> kubectl port-forward svc/argocd-server -n argocd 8080:443
- get argocd password -> kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath='{.data.password}' | base64 --decode

## Deploy application to argocd
- k apply -f k8s/app.argo.yaml


Extra reading: [here](https://www.arthurkoziel.com/setting-up-argocd-with-helm/)