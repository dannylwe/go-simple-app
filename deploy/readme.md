## create kindl cluster
```bash
kind create cluster --name dev-observability --config monitoring/kind-cluster.yaml
```

## install argocd
```bash
helm install argocd argo/argo-cd -n argocd --create-namespace -f deploy/helm/argocd/values.yaml --debug --wait --timeout 20m
```
- get admin password: `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d`

## install helm charts for monitoring
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add loki https://grafana.github.io/helm-charts
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
```

## install pv and pvc
```bash
kubectl apply -f https://raw.githubusercontent.com/rancher/local-path-provisioner/master/deploy/local-path-storage.yaml
```
## grafana
```bash
helm install prometheus prometheus-community/kube-prometheus-stack   -n monitoring   -f deploy/helm/prometheus/values.yaml --debug --wait --timeout 10m
```

```bash
kubectl --namespace monitoring get secrets prometheus-grafana -o jsonpath="{.data.admin-password}" | base64 -d ; echo
```

## grafana
```bash
helm install grafana grafana/grafana \
  -n monitoring \
  -f deploy/helm/grafana/values.yaml --debug --wait --atomic --timeout 10m
```

```bash
kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
```

## loki
```bash
helm install loki grafana/loki-stack \
  -n monitoring \
  -f deploy/helm/loki/values.yaml --debug --wait --atomic --timeout 10m
```

## prom-tail
```bash
helm install promtail grafana/promtail \
  -n monitoring \
  -f deploy/helm/promtail/values.yaml --debug --wait --timeout 10m
```
## otel-collector (failed to install)
```bash
helm install otel-collector open-telemetry/opentelemetry-collector \
  -n monitoring \
  -f deploy/helm/otel-collector/values.yaml --debug --wait --timeout 10m
```

## tempo
```bash
helm install tempo grafana/tempo-distributed \
  -n monitoring \
  -f deploy/helm/tempo/values.yaml --debug --timeout 10m
```

====================
Service DNS
====================
http://<service_name>.<namespace>.svc.cluster.local:<port>