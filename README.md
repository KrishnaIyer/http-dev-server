# http-dev-server

A simple HTTP development server written in Golang.

## Usage

Run directly from docker.

```bash
$ docker run --rm -p 8080:8080 krishnaiyer/http-dev-server
```

## Helm

1. Setup a Kubernetes Cluster.

2. Install the Traefik CRDs.

```bash
$ kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.9/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
```

## License

The contents of this repository are provided `as-is` under the terms of the [Apache 2.0](./LICENSE) license.
