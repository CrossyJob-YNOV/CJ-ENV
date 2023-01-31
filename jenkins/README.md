# Jenkins

## Installation

1. Make a symbolic link from `/data/CJ-ENV/jenkins/` to `/home/jenkins/`
2. Copy the `docker-jenkins.service` file to `/etc/systemd/system/`
3. Run `systemctl daemon-reload`
4. Enable the unit with `systemctl enable docker-jenkins`

Then, add the jobs through the Jenkins UI, which is running on **/jenkins**.
