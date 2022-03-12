
#!/bin/sh

HERE=$(cd $(dirname $BASH_SOURCE) && pwd)

sh ${HERE}/clear.sh
docker run \
    --name="consul-helm-teller" \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


until $(curl --output /dev/null --silent --head --fail http://127.0.0.1:8500); do
    echo 'Waiting for consul service to be ready'
    sleep 5
done


echo "adding redis config"
docker exec consul-helm-teller sh -c "consul kv put redis/config/host localhost"
docker exec consul-helm-teller sh -c "consul kv put redis/config/password 1234"
echo "adding log level config"
docker exec consul-helm-teller sh -c "consul kv put log-level debug"

echo "You can run clear.sh for cleanup"