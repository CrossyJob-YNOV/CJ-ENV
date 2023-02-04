FROM jenkins/jenkins:jdk17

USER root

RUN apt update && apt upgrade -y && apt install -y binutils xz-utils gcc maven apt-transport-https ca-certificates curl gnupg2 software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" && \
    apt update && apt install -y docker-ce docker-ce-cli containerd.io && \
    apt clean all

RUN usermod -aG docker jenkins
