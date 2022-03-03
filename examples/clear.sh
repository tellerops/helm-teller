
#!/bin/sh

echo "stop consul"
docker stop consul-helm-teller
echo "remove container name"
docker container rm consul-helm-teller