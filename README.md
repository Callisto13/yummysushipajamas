[![Build Status](https://travis-ci.org/Callisto13/yummysushipajamas.svg?branch=master)](https://travis-ci.org/Callisto13/yummysushipajamas)

### What is it?

A very basic gRPC server and CLI based client.

It is designed primarily to run in containers/be cloud native. Kube manifests
are included.

Obviously everything has been TDD-ed, see below for running tests. (Ensure [ginkgo](https://onsi.github.io/ginkgo/) is installed first,
or just run everything is Docker, which is fine but slower.) 

Deps are managed with `go mod`, and everything has been built and tested
with Go version 1.13.3.

### What does it do?

Very little and nothing useful :)

```sh
# sum 2 numbers
$ ./bin/client -action=sum 3 5
8
# list all primes between 2 numbers (hopefully, implementation is very naive atm)
$ ./bin/client -action=prime 0 20
2
3
5
7
11
13
17
19
```

### What _could_ it do?

When I was nearly done I realised that this could be more useful (or useful at all), so on another branch
I am playing with turning this into a tool which can debug processes in a Pod.
Perhaps return interesting things like namespace ids and cgroup configuration, as well
as more useful stuff like process states.

Of course that would mean deploying Pods with [shared PID namespaces](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/)
otherwise there would be very little to see. PaaSes like CF (and AFAIK Heroku) put their
siecars in shared PID NSes by default so that should just work. On regular Kube deployments there
may not be much need for this, as they are not multi-tenant so most devs would have access
to container (even host) cmdline. But on multi-tenant PaaSes who, for whatever reason, have _not_ enabled ssh into
apps, it may be cool? ¯\\\_(ツ)\_/¯

### Deploying yummysushipajamas locally with minikube

**Prerequisites:**
- [`ytt`](https://get-ytt.io/)
- [`minikube`](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [`docker`](https://docs.docker.com/v17.09/engine/installation/)

**Then:**
1. Start the cluster: `minikube start`.
1. Clone this repo and `cd` into it.
1. Deploy the `ysp-server` and the associated service: `make deploy`.
1. Copy and paste the suggested command to set the server address env var.
1. Call the client: `./bin/client`.

**Configuration options:**

- The server port can be edited in `kube/values.yaml`

### Building the server and client and doing other things
```sh
$ make help
Usage:
  proto   ..................... regenerate grpc sources (will go to ./ysp/ysp.proto)
  client  ..................... build the client bin (will go to ./bin/client)
  server  ..................... build the server bin (will go to ./bin/server)
  clean   ..................... delete bins
  test    ..................... run all test suites in docker container
  unit    ..................... run server and client unit tests
  int     ..................... run integration tests in docker container
  dep     ..................... update dependencies
  mock    ..................... regenerate grpc testing mocks
  docker  ..................... rebuild and push callisto13/ysp docker image
  deploy  ..................... deploy server and service to minikube
  destroy ..................... delete server and service
  reload  ..................... proto server client docker destroy deploy (aka rebuild and redeploy the lot)
```

### Repo structure
```
.
├── Dockerfile
├── Makefile
├── README.md
├── bin
│   ├── client
│   └── server
├── ci
│   └── Dockerfile
├── client
│   ├── client.go
│   ├── client_suite_test.go
│   ├── client_test.go
│   └── cmd
│       └── main.go
├── go.mod
├── go.sum
├── integration
│   ├── integration_suite_test.go
│   └── integration_test.go
├── kube
│   ├── deployment.yaml
│   ├── service.yaml
│   └── values.yaml
├── pb
│   ├── mocks
│   │   └── ysp_mock.go
│   ├── ysp.pb.go
│   └── ysp.proto
├── server
│   ├── cmd
│   │   └── main.go
│   ├── server.go
│   ├── server_suite_test.go
│   └── server_test.go
└── vendor
    └── ...
```
