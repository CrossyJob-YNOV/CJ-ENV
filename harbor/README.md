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

5. Comme nous n'utilisons pas HTTPS, on ajoute notre registry à la liste des **insecure registries** :
   ```shell
   vi /etc/docker/daemon.json
   ```
   Ajouter à la fin :
   ```json
   "insecure-registries": ["51.159.160.76:8081"]
   ```


## Ajout du SSL
Dans un second temps, comme github ne pouvait push des images dockers sur des registry non sécurisé, nous avons généré des certificats SSL avec certbot.

De ce fait, on re modifie la configuration d'harbor.

```
hostname: harbor.crossyjob.ezyostudio.com

# http related config
http:
  # port for http, default is 80. If https enabled, this port will redirect to https port
  port: 8081

# https related config
#https:
  # https port for harbor, default is 443
  #port: 8082
  # The path of cert and key files for nginx
 # certificate: /home/harbor/certif/fullchain.pem
  #private_key: /home/harbor/certif/privkey.pem
```

On ajoute donc le hostname avec notre url d'accès, comme le HTTPS est géré par nginx, on désactive la configuration en https avec harbor. Le reverse proxy d'nginx se fera entre le port 443 du server et le port 8081 du container d'harbor.

```
sudo certbot --nginx -d example.com -d www.example.com 


Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/crossyjob.ezyostudio.com/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/crossyjob.ezyostudio.com/privkey.pem
This certificate expires on 2023-05-08.
These files will be updated when the certificate renews.
Certbot has set up a scheduled task to automatically renew this certificate in the background.

Deploying certificate
Successfully deployed certificate for crossyjob.ezyostudio.com to /etc/nginx/sites-enabled/nginx.conf
Congratulations! You have successfully enabled HTTPS on https://crossyjob.ezyostudio.com
```

Les certificats sont générés, on configure maintenant nginx pour nous rediriger vers harbor.

```
server {
    server_name harbor.crossyjob.ezyostudio.com; # managed by Certbot


    location / {
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
        proxy_set_header Host $host;
        proxy_set_header Origin http://$host;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Location $scheme$host$request_uri;
        proxy_pass http://127.0.0.1:8081/;
    }


    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/harbor.crossyjob.ezyostudio.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/harbor.crossyjob.ezyostudio.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
```

Ici le serveur nginx écoute le port 443 avec un host provenant du host de notre certificat (harbor.crossyjob.ezyostudio.com) et redirige vers le port du docker harbor.