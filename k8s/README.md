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

3. Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

4. Download Minikube:
    ```shell
    curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
      && chmod +x minikube
    ```

5. Install Minikube:
    ```shell
    sudo mkdir -p /usr/local/bin/
    sudo install minikube /usr/local/bin/
    ```
   
6. Start Minikube:
    ```shell
    systemctl start minikube-kubernetes
    ```
   
7. Enable Ingress
   ```shell
   minikube addons enable ingress
   minikube addons enable ingress-dns
   ```
