# Intégration continue github action

Ci joint le dockerfile pour build l'app nuxt, ainsi que la pipeline github.
Lorsqu'on merge sur main, la pipeline se lance, build une image docker de l'app et (reste a faire) la push sur le registry harbor.