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
- run postgres database `docker run --rm --name foo-postgres -d -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -p 5432:5432 postgres:16.2`
- *(optional)* run telemetry exporter `docker run --rm --name foo-otel-collector -p 4317:4317 otel/opentelemetry-collector-contrib:0.85.0`

### Running locally
- run server `make run-server`
- run worker `make run-worker`
