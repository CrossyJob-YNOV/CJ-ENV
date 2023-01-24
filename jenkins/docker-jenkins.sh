#!/bin/bash -eu

COMPOSE_PATH="$HOME/jenkins/jenkins-docker-compose.yml"

checkError() {
    if [[ $# -eq 2 ]] && [[ $1 -ne 0 ]]; then
        echo "Failed to start $2"
        exit 1
    fi
}

dockerComposeUp() {
    echo "Composing Jenkins containers ..."

    docker-compose -f $COMPOSE_PATH up -d
    checkError $? "Jenkins compose up"

    echo "Successfully composed Jenkins containers !"
}

dockerComposeDown() {
    echo "Stopping Jenkins containers ..."
    docker-compose -f $COMPOSE_PATH down
    checkError $? "Jenkins compose down"

    echo "Successfully stopped Jenkins containers !"
}

if [[ $# -eq 0 ]] || [[ $1 == "start" ]]; then
    dockerComposeUp
elif [[ $1 == "stop" ]]; then
    dockerComposeDown
else
    echo "Invalid syntax, use $0 [start|stop]"
    exit 1
fi
