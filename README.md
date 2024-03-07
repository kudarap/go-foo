# Foo Service
golang backend service scaffold

### Features
- [x] API server
- [x] worker for background processing or consumers
- [x] telemetry
- [x] automatic database migration
- [x] config file with env var override
- [x] unit tests

### Requirements
- go 1.22
- docker 24

### Setup
- copy `.env.sample` to `.env` and change values accordingly
- run postgres database `make local-dbs`
- *(optional)* run telemetry exporter `make local-otel-collector`

### Running locally
- run server `make run-server`
- run worker `make run-worker`
