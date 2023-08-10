#!/bin/bash
# how to use: ./redeploy.sh {image_tag} in CLI
docker stop snippetbox
docker rm snippetbox
docker image rm {image_url}:$1
docker run -d -m 2048M --name snippetbox \
  --publish {available_port}:{container_port} --network={available_network} --ip={available_ip} \
  {image_url}:$1

docker network connect --ip={another_available_ip} {another_available_network} snippetbox

# Optional: copy properties file / auth key to container
docker cp /opt/app-properties/snippetbox.properties snippetbox:/opt/application.properties
docker cp /opt/public_key.pem snippetbox:/opt/public_key.pem

