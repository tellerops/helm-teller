<p align="center">
<br/>
<br/>
<br/>
   <img src="media/helm-teller.svg" width="395"/>
<br/>
<br/>
</p>

<p align="center">
<b>:computer: Never leave your terminal for secrets</b>
<br/>
<b>:pager: Same workflows for all your environments</b>
<br/><br/><br/>
<hr/>
</p>

# Helm-teller

Helm [Teller](https://github.com/SpectralOps/teller) allows you to pull secrets or configuration setting when you deploy a helm chart.

## Why should i use it?
* More secure while using `--debug` or `--dry-run` the secrets will not showed in the STDOUT
* Simple to integrate
* Rich of supported plugins
* Pull configuration and secret from multiple provider in one place
* Manage configuration from development to production in same way


![](media/helm-teller.gif)

## Installation
```sh
$ helm plugin install https://github.com/SpectralOps/helm-teller
```

## Quick Start with helm teller
* Create [.teller.yaml](https://github.com/SpectralOps/teller#quick-start-with-teller-or-tlr) file in your helm chart.
```yaml
providers:
  # vault provider
  vault:
    env_sync:
      path: redis/config
  # Consul provider
  consul:
    env:
      loglevel:
        path: log-level
 
```
* Set teller fields in your helm chart 
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config-map
data:
  redis-host: {{ .Values.teller.host }}
  redis-password: {{ .Values.teller.password }}
  loglevel: {{ .Values.teller.loglevel }}
```
* Run helm teller deploy `helm teller [install/upgrade] -- {HELM FLAGS}`.
 

See working example [here](./examples/examples.md)

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