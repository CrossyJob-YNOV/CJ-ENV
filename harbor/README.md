# Harbor

## Installation

1. On récupère la derniere release sur github https://github.com/goharbor/harbor/releases/latest

2. On la décompresse avec
    ```shell
    tar xzvf harbor-offline-installer-v2.7.0.tgz 
    ```

3. Dans le dossier harbor, on récupère le fichier de configuration *harbor.yml.tmpl*
    - On y renseigne le hostname, ici l'ip externe de notre server, ansi que le port sur lequel on exécutera harbor.

4. On lance le programme d'installation, puis une fois terminé, on retrouve harbor dans notre navigateur
   à l'adresse du hostname.
