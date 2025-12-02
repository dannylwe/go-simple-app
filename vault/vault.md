## install mkcert
- brew install mkcert nss

## install local ca
- mkcert -install
- mkcert vault.vault.svc.cluster.local vault.vault.svc vault

mv vault.vault.svc.cluster.local+2-key.pem vault-key.pem
mv vault.vault.svc.cluster.local+2.pem vault.pem
cp vault.pem vault.crt
cp vault-key.pem vault.key

## create namespace and create tls secret
- kubectl create ns vault || true
- kubectl create secret tls vault-server-tls --cert=vault.crt --key=vault.key -n vault

## values.yaml for vault helm
```bash
# vault-values.yaml
server:
  image:
    repository: hashicorp/vault
    tag: "1.14.0" # pick a stable vault version; adjust if newer desired
  ha:
    enabled: true
    replicas: 1
  # Enable integrated raft storage (recommended for production-like)
  dataStorage:
    enabled: true
    volumeSize: 500Mi
    storageClassName: ""
    # storageClassName: "standard" # leave empty to use default minikube SC
    raft:
      enabled: true

  # TLS configuration: use an existing k8s TLS secret
  tls:
    enabled: false
    # existingSecret: "vault-server-tls"
    # The server will use /etc/vault/tls/tls.crt and tls.key if you used secret tls type

  # enable the UI
  ui:
    enabled: true

  # enable the injector (mutating webhook and vault-agent)
  injector:
    enabled: true

  # service config
  service:
    type: ClusterIP
    port: 8200
    annotations: {}

  # resources for each vault server pod (tweak for minikube limits)
  resources:
    requests:
      cpu: 100m
      memory: 256Mi
    limits:
      cpu: 500m
      memory: 512Mi

# Enable server-side RBAC and service account (default)
rbac:
  create: true

# TLS on the client side (for UI/CLI access)
serverConfig: |
  listener "tcp" {
    address     = "0.0.0.0:8200"
    tls_cert_file = "/vault/userconfig/vault.crt"
    tls_key_file  = "/vault/userconfig/vault.key"
  }

# Use PostStart hook to ensure proper permissions, etc.
affinity: {}
tolerations: []
```

## helm install
```bash
helm upgrade --install vault hashicorp/vault \
  --namespace vault \
  -f vault-values.yaml \
  --timeout 10m
  --wait
```

## upgrade helm
```bash
helm upgrade vault hashicorp/vault --namespace vault -f vault-values.yaml
```
TBD production


## run in dev mode
```zsh
helm install vault hashicorp/vault \
       --set='server.dev.enabled=true' \
       --set='ui.enabled=true' \
       --set='ui.serviceType=LoadBalancer' \
       --namespace vault
```

## create policy
```zsh
kubectl exec -it vault-0 -n vault -- /bin/sh
```
```shell
cat <<EOF > /home/vault/read-policy.hcl
path "secret*" {
  capabilities = ["read"]
}
EOF
```

## Add policy
```zsh
vault policy write read-policy /home/vault/read-policy.hcl
```

## Add kubernetes auth
```bash
vault auth enable kubernetes
```

## Configure Vault to communicate with the Kubernetes API server
```bash
vault write auth/kubernetes/config \
   token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
   kubernetes_host=https://${KUBERNETES_PORT_443_TCP_ADDR}:443 \
   kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```

## create role
- create role (vault-role) that binds the created policy to a service account (vault-serviceaccount)
in a specific namespace
NOTE: To add multiple namespaces and/or service accounts, they need to be comma separated

```bash
vault write auth/kubernetes/role/vault-role \
   bound_service_account_names=vault-serviceaccount \
   bound_service_account_namespaces=vault \
   policies=read-policy \
   ttl=10h
#    token_audiences="vault"
```

## create service account
```bash
cat <<EOF > vault-sa.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-serviceaccount
  labels:
    app: read-vault-secret
EOF
```

## enable kv v2
```bash
vault secrets enable -path=secret kv-v2
```

## put in vault kv
```bash
vault kv put secret/login pattoken=ytbuytbytbf765rb65u56rv
```

`vault kv list secret`
## port forward
```bash
kubectl port-forward svc/vault 8200:8200 -n vault
```

## stop port forwarding
```bash
jobs -l; kill $jobId
```

## secrets location
```bash
k exec -it <podname> -n <namespace> -- /bin/sh
```