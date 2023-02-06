# Kubernetes (Minikube)

## Installation

1. ```shell
   ln -s /data/CJ-ENV/k8s/ /home/kubernetes/
   ```

2. ```shell
   ln -s /data/CJ-ENV/k8s/minikube-kubernetes.service /etc/systemd/system/
   systemctl daemon-reload
   systemctl enable minikube-kubernetes
   ```

3. Download [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
   ```shell
   curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
   ```

4. Install kubectl
   ```shell
   sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
   ```

5. Download Minikube:
    ```shell
    curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
      && chmod +x minikube
    ```

6. Install Minikube:
    ```shell
    sudo mkdir -p /usr/local/bin/
    sudo install minikube /usr/local/bin/
    ```

7. Start Minikube:
    ```shell
    systemctl start minikube-kubernetes
    ```

8. Enable Ingress
   ```shell
   minikube addons enable ingress
   minikube addons enable ingress-dns
   ```

9. Create all deployments, services, ...
   ```shell
   cd ~kubernetes/k8s
   kubectl apply -f back/
   kubectl apply -f front/
   kubectl apply -f ci-listener/
   ```
