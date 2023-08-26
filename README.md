# Api TrxReceiver (Transaction Receiver)

## How-To

### Run

`API_PORT=4000 go run main.go`

### Call the APIs

- GET `curl localhost:4000/`
- GET `curl localhost:4000/health`
- GET `curl localhost:4000/trx`
- POST `curl -X POST localhost:4000/trx -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'`
- PUT `curl -X PUT localhost:4000/trx/123 -d '{"id": 123, "symbol": "CSS", "name": "Counter-Strike Source", "Value": 9001}'`

