# Create Cluster
gcloud container clusters create "hubble-test-03" \
  --node-taints node.cilium.io/agent-not-ready=true:NoExecute \
  --zone us-west2-a \
  --num-nodes 3 \
  --disk-size=50gb

# Install Prometheus
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
kubectl create ns monitoring
helm install kps -n monitoring prometheus-community/kube-prometheus-stack -f ./kps-values.yml

# Install Cilium
CILIUM_CLI_VERSION=$(curl -s https://raw.githubusercontent.com/cilium/cilium-cli/main/stable.txt)
CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/cilium/cilium-cli/releases/download/${CILIUM_CLI_VERSION}/cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check cilium-linux-${CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC cilium-linux-${CLI_ARCH}.tar.gz /usr/local/bin
rm cilium-linux-${CLI_ARCH}.tar.gz{,.sha256sum}

cilium install --chart-directory ./cilium/install/kubernetes/cilium \
  --namespace kube-system \
  --set prometheus.enabled=true \
  --set operator.prometheus.enabled=true \
  --set hubble.enabled=true \
  --set hubble.metrics.enableOpenMetrics=true \
  --set hubble.metrics.enabled="{dns,drop,tcp,flow:labelsContext=source_ip\,source_namespace\,source_workload\,destination_ip\,destination_namespace\,destination_workload\,traffic_direction,flows-to-world,port-distribution,icmp,httpV2:exemplars=true;labelsContext=source_ip\,source_namespace\,source_workload\,destination_ip\,destination_namespace\,destination_workload\,traffic_direction}" \
  --set hubble.metrics.serviceMonitor.enabled=true \
  --set hubble.ui.enabled=true \
  --set hubble.relay.enabled=true \
  --set gke.enabled=true --dry-run-helm-values

# Sample application and policy
```
apiVersion: v1
kind: Service
metadata:
  name: frontend
spec:
  type: ClusterIP
  ports:
    - port: 80
  selector:
    app: frontend
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: frontend
  template:
    metadata:
      labels:
        app: frontend
        env: prod
    spec:
      containers:
        - name: nginx
          image: ubuntu/nginx:latest
          ports:
            - containerPort: 80 
---
apiVersion: v1
kind: Service
metadata:
  name: backend
spec:
  type: ClusterIP
  ports:
    - port: 80
  selector:
    app: backend
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend
spec:
  replicas: 1
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
        env: prod
    spec:
      containers:
        - name: nginx
          image: ubuntu/nginx:latest
          ports:
            - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: db
spec:
  type: ClusterIP
  ports:
    - port: 80
  selector:
    app: db
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: db
spec:
  replicas: 1
  selector:
    matchLabels:
      app: db
  template:
    metadata:
      labels:
        app: db
        env: prod
    spec:
      containers:
        - name: nginx
          image: ubuntu/nginx:latest
          ports:
            - containerPort: 80
---
apiVersion: "cilium.io/v2"
kind: CiliumNetworkPolicy
metadata:
  name: "backend-db-rule"
spec:
  description: "policy to restrict access to database call"
  endpointSelector:
    matchLabels:
      app: db
  ingress:
  - fromEndpoints:
    - matchLabels:
        app: backend
        env: prod
    toPorts: 
    - ports:
      - port: "80"
        protocol: TCP
      rules:
        http:
        - method: "GET"
          path: "/"
```

# Sample grafana dashboards to import
19423
19424
19425

# Generate load
apt update && apt -y install curl wget
wget https://hey-release.s3.us-east-2.amazonaws.com/hey_linux_amd64
chmod +x ./hey_linux_amd64

./hey_linux_amd64 -c 1 -q 10 -n 10000000 -t 1 http://db




kubectl annotate pod -n default deathstar-7848d6c4d5-kj7t4 policy.cilium.io/proxy-visibility="<Egress/53/UDP/DNS>,<Egress/80/TCP/HTTP>" --overwrite



kubectl create -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/minikube/http-sw-app.yaml
kubectl apply -f https://raw.githubusercontent.com/cilium/cilium/HEAD/examples/minikube/sw_l3_l4_l7_policy.yaml


./hey_linux_amd64 -c 5 -q 200 -n 100000 -m POST http://deathstar.default.svc.cluster.local/v1/request-landing


a

./hey_linux_amd64 -c 5 -q 200 -n 100000 -m POST http://deathstar.default.svc.cluster.local/v1/request-landing

./hey_linux_amd64 -c 5 -q 200 -n 100000 -t 1 http://frontend
./hey_linux_amd64 -c 5 -q 200 -n 100000 -t 1 http://backend
./hey_linux_amd64 -c 5 -q 200 -n 100000 -t 1 http://db

./hey_linux_amd64 -c 10 -q 1000 -n 10000 -t 1 http://db/restricted
./hey_linux_amd64 -c 1 -q 10 -n 10000000 -t 1 http://db
./hey_linux_amd64 -c 1 -q 100 -n 10000000 -t 1 http://db



cat << EOF > parseable-env-secret
addr=0.0.0.0:8000
staging.dir=./staging
fs.dir=./data
username=admin
password=admin
EOF

kubectl create ns parseable
kubectl create secret generic parseable-env-secret --from-env-file=parseable-env-secret -n parseable

helm repo add parseable https://charts.parseable.io
helm install parseable parseable/parseable -n parseable --set "parseable.local=true"
kubectl port-forward svc/parseable 8000:80 -n parseable

helm repo add vector https://helm.vector.dev


curl --cacert /etc/srv/kubernetes/pki/ca-certificates.crt --header "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImZhQmtRaG85UzV0QTBiYVl4TGxBYkN3UjRoNkp1MkRTWm1XNmJPTEFCY2cifQ.eyJhdWQiOlsiaHR0cHM6Ly9jb250YWluZXIuZ29vZ2xlYXBpcy5jb20vdjEvcHJvamVjdHMvcmwtbW9yZG9yL2xvY2F0aW9ucy91cy13ZXN0Mi1hL2NsdXN0ZXJzL2h1YmJsZS10ZXN0LTAzIl0sImV4cCI6MTcyOTE2MjMwNSwiaWF0IjoxNjk3NjI2MzA1LCJpc3MiOiJodHRwczovL2NvbnRhaW5lci5nb29nbGVhcGlzLmNvbS92MS9wcm9qZWN0cy9ybC1tb3Jkb3IvbG9jYXRpb25zL3VzLXdlc3QyLWEvY2x1c3RlcnMvaHViYmxlLXRlc3QtMDMiLCJrdWJlcm5ldGVzLmlvIjp7Im5hbWVzcGFjZSI6ImRlZmF1bHQiLCJwb2QiOnsibmFtZSI6ImV2ZXJ5dGhpbmctYWxsb3dlZC1leGVjLXBvZCIsInVpZCI6Ijc0NjBkYzVmLWE2NWYtNDE2YS04MDJmLTNkNjAxMDY3YjhjMCJ9LCJzZXJ2aWNlYWNjb3VudCI6eyJuYW1lIjoiZGVmYXVsdCIsInVpZCI6IjM4MDUxODE1LTZjMTktNGI2ZC04NjM2LTFkN2JkOTFhYjY5ZCJ9LCJ3YXJuYWZ0ZXIiOjE2OTc2Mjk5MTJ9LCJuYmYiOjE2OTc2MjYzMDUsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.c4iTc1vwH4H25g1Du9Qc0yXvE-QcLljRxnA-VX3kexlVTUrDdraox_5lFeU3TSglHKdAiD9m157hAGnR2Qgu_z_PfcEemNYoo_UcydnhiQ3zdayLx1OHSAkMM7X-zcxGU4l5CQ2XOg85L3jvespbXOfbpf8S6yHE-S7nZ1B5atEe9F1uJNHcp6Nnb0esO_BXkgvVFTfg2RbwuBzahD_Lu6vjs-IEvKXMNfRPvIJCiLHQJkt38jlKqjIxjKl0VLUy_pZYEC3WjIU_SaT1ZXHbhpg3KyaVmdoiItF5yXXzma0qyMzLBIqN2KNg2qP7JOlxdoeB5LkzQj4G8NsUnOsHvgroot" -X GET https://10.168.0.28/api
