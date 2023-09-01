# Api TrxReceiver (Transaction Receiver)

## How-To

To build and run applications (locally), please install Rancher Desktop or another container orchestration tool (Rancher, Docker, etc.)
For the purpose of this README, we are using Rancher Desktop

### Build using Paketo

#### Prerequisites

- Download BuildPacks CLI by following these instructions here: https://buildpacks.io/docs/tools/pack/
- Expose the docker socket in the DOCKER_HOST env var
    ```shell
    export DOCKER_HOST="unix://${HOME}/.rd/docker.sock"
    ```

#### Build using Google's builder

- Create the image using the buildpack:
  ```shell
  pack build trxreceiver-app --path . --builder gcr.io/buildpacks/builder:v1
  v1: Pulling from buildpacks/builder
  53e5e158da5a: Pulling fs layer
  7577c9c60d3f: Pulling fs layer
  3c2cba919283: Pulling fs layer
  29c0b7dfe559: Pulling fs layer
  461089a250f0: Pulling fs layer
  3e8e410021b1: Pulling fs layer
  b1c668cf7c45: Pulling fs layer
  ... (more output here)
  Adding cache layer 'google.go.gomod:gopath'
  Timer: Cache ran for 2.173747273s and ended at 2023-09-01T01:42:17Z
  Successfully built image trxreceiver-app
    ```

### Local run

#### Run Redis:
```shell

❯ docker run --name redis-db -p 127.0.0.1:6379:6379/tcp -d docker.io/redis
Trying to pull docker.io/library/rdb:latest...
Getting image source signatures
Copying blob 29771da5b50b done  
Copying blob 16acd9ca1349 done  
Copying blob 723b2c9888ad done  
Copying blob 2f34c7846499 done  
Copying blob 52d2b7f179e3 done  
Copying blob 689bed60e397 done  
Copying config 506734eb5e done  
Writing manifest to image destination
Storing signatures
3e62aa16fe25225caf279e5abee810f41c5fdec28b9a831ad70abf2f31747fd2

❯ podman container ls
CONTAINER ID  IMAGE                           COMMAND       CREATED             STATUS                 PORTS       NAMES
3e62aa16fe25  docker.io/library/rdb:latest  rdb-server  About a minute ago  Up About a minute ago              rdb-db
```

#### Run the API

Using the container that the buildpack built

```shell
docker run --name trxreceiver-app -p 127.0.0.1:4000:4000/tcp --env API_PORT=4000 --env REDIS_HOST=localhost --REDIS_PORT=6379 -d trxreceiver-app
19b1e6445d05726a7809c3453456761accf23c0477976a0b38ae779d2504f91e
❯ docker container ls
CONTAINER ID   IMAGE             COMMAND              CREATED         STATUS         PORTS                      NAMES
19b1e6445d05   trxreceiver-app   "/cnb/process/web"   5 seconds ago   Up 2 seconds   127.0.0.1:4000->4000/tcp   trxreceiver-app
```

Or running locally using go

Run this App/APIs
```shell
export API_PORT=4000 REDIS_HOST=localhost REDIS_PORT=6379
go run main.go
```

### Run Unit tests

```shell
go test ./...
```

### Kubernetes

TBD

## API Calls

| Operation      | cURL example                                                                                                           |
|:---------------|:-----------------------------------------------------------------------------------------------------------------------|
| GET root       | `curl localhost:4000/`                                                                                                 |
| GET api health | `curl localhost:4000/health`                                                                                           |
| GET trx        | `curl -X GET localhost:4000/trx/123`                                                                                   |
| POST trx       | `curl -X POST localhost:4000/trx -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'`    |
| PUT trx        | `curl -X PUT localhost:4000/trx/123 -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'` |

## Next

### Done
1. Create a simple API using go-chi with GET POST PUT ops
2. Add logging
3. Persist changes into an in-memory Redis cache
4. Unit testing in go
5. Package app using buildpacks: https://buildpacks.io/docs/app-developer-guide/build-an-app/

### TODO
6. Implement k6 for load testing on existing k8s cluster
7. Add k8s manifests deployments here
8. View metrics on Prometheus + Grafana