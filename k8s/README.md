# Kubernetes (Minikube)

## Installation

1. Make a symbolic link from `/data/CJ-ENV/k8s/` to `/home/kubernetes/`
2. Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
3. Download Minikube:
    ```shell
    curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
      && chmod +x minikube
    ```
4. Install Minikube:
    ```shell
    sudo mkdir -p /usr/local/bin/
    sudo install minikube /usr/local/bin/
    ```
5. Start Minikube:
    ```shell
    minikube start --driver=docker
    ```
