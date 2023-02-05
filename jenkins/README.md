# Jenkins

## Installation

1. ```shell
    ln -s /data/CJ-ENV/jenkins/ /home/jenkins/
    ```
   
2. ```shell
    ln -s /data/CJ-ENV/jenkins/docker-jenkins.service /etc/systemd/system/
    systemctl daemon-reload
    systemctl enable --now docker-jenkins
    ```

Then, add the jobs through the Jenkins UI, which is running on **/jenkins**.
