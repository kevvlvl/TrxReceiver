# Api TrxReceiver (Transaction Receiver)

## How-To

### Local run

To run locally, please install podman or another container orchestration tool (Rancher, Docker, etc.)
For the purpose of this README, we are using podman

Run Redis:
```shell

❯ podman run --name redis-db -p 127.0.0.1:6379:6379/tcp -d docker.io/redis
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

Run this App/APIs
```shell
API_PORT=4000 go run main.go
```

### Kubernetes

## API Calls

| Operation      | cURL example                                                                                                           |
|:---------------|:-----------------------------------------------------------------------------------------------------------------------|
| GET root       | `curl localhost:4000/`                                                                                                 |
| GET api health | `curl localhost:4000/health`                                                                                           |
| GET stock trx  | `curl -X GET localhost:4000/trx/123`                                                                                   |
| POST stock trx | `curl -X POST localhost:4000/trx -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'`    |
| PUT stock trx  | `curl -X PUT localhost:4000/trx/123 -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'` |

## Next

1. Create a simple API using go-chi with GET POST PUT ops - DONE
2. Add logging - DONE
3. Persist changes into an in-memory Redis cache - DONE
4. Unit testing in go - NEXT
5. Package app using Paketo buildpacks: https://paketo.io/docs/howto/go/
6. Implement k6 for load testing on existing k8s cluster
7. Add k8s manifests deployments & Paketo buildpacks here
7View metrics on Prometheus + Grafana