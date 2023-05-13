#!/bin/sh

cd ../
DTAG=$(date +"%Y%m%d%H%M%S")
docker build --platform linux/amd64 -f asgard-gateway/Dockerfile -t asgard-gateway .
docker tag asgard-gateway registry.digitalocean.com/francisco/asgard-gateway:$DTAG
docker push registry.digitalocean.com/francisco/asgard-gateway:$DTAG

echo "registry.digitalocean.com/francisco/asgard-gateway:$DTAG image push success"