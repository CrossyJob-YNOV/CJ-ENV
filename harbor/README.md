Procédure d'installation d'harbor 

1. On récupère la derniere release sur github https://github.com/goharbor/harbor/releases

2. On la décompresse avec 
```
tar xzvf harbor-offline-installer-v2.7.0.tgz 
```

3. Dans le dossier harbor, on récupère le fichier de configuration *harbor.yml.tmpl*
- On y reseigne l'hostname, ici l'ip externe de notre server, ansi que le port sur lequel on executera harbor.

4. On lance le programme d'installation, une fois terminé, on retrouve harbor dans notre navigateur à l'adresse de l'hostname.



