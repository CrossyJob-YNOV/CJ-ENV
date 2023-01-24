FROM jenkins/jenkins:jdk17

USER root

RUN apt update && apt upgrade -y && apt install -y binutils xz-utils gcc
