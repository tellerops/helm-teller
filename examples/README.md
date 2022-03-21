# Example

## Pre-requirements
* helm-teller installed
* running Kubernetes cluster 

Run [run.sh](./run.sh) file of do the following steps. run [clear.sh](./clear.sh) to delete `docker` container after your test.

## You can create testing environment by follow the steps:

### Run consul provider
```sh
$ docker run \
    --name="consul-helm-teller" \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

### Set keys
```sh
$ docker exec consul-helm-teller sh -c "consul kv put redis/config/host localhost"
$ docker exec consul-helm-teller sh -c "consul kv put redis/config/password 1234"
$ docker exec consul-helm-teller sh -c "consul kv put log-level debug"
```

## Install the chart
```sh
$ cd examples/example-chart
$ helm teller install -- --debug --dry-run  test .
```

Delete consul container
```sh
$ docker stop consul-helm-teller
$ docker container rm consul-helm-teller
```