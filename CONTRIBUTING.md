# Contributing


## Developing
1. Make sure you have configured running Kubernetes cluster.
```sh
$ kubectl cluster-info
```
2. configure [.teller.yaml](./examples/example-chart/.teller.yaml) to pull variables from your various cloud providers. you can also run [run.sh](./examples/run.sh) to work via local Consul.
3. Run local helm-teller via [example-chart](./examples/example-chart/).
```sh
$ go run main.go install --teller-config "examples/example-chart/.teller.yaml" -- --debug --dry-run  test ./examples/example-chart
```

## Testing
To run unit tests:
```sh
$ make test
```

### Linting
Linting is treated as a form of testing (using `golangci`), to run:
```sh
$ make lint
```

### Formatting
Run go fmt check:
```sh
$ make fmt
```

### Run all checks:
```sh
$ make checks
```