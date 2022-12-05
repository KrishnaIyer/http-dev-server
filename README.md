# http-dev-server

A simple HTTP development server written in Golang.

## Usage

Run directly from docker.

```bash
$ docker run --rm -p 8080:8080 krishnaiyer/http-dev-server
```

## Helm

> This requires Helm v3.8.0 and above.

1. Setup a Kubernetes Cluster.

2. Install the Traefik CRDs.

```bash
$ kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.9/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
```

3. Pull the Helm Charts

```bash
$ helm pull oci://registry-1.docker.io/krishnaiyer/http-dev-server-helm --version <version>
```

## License

The contents of this repository are provided `as-is` under the terms of the [Apache 2.0](./LICENSE) license.
