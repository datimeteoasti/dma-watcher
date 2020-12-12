# dma-watcher

## Build

```bash
make build
```

## Configure

Before running this project on your local development environment, you need to configure the credentials that dma-watcher uses to connect to the database.

`dma-watcher` reads the connection string from the environment variable `DATABASE_URL`, so you need to export it in the following way

```bash
export DATABASE_URL="postgresql://postgres:password@localhost:5432/data"
```

To run a PostgreSQL database locally, you can leverage the docker-compose file, which defines a service called `DB`. To spin up a new instance, make sure that your Docker daemon is running and run the following command:

```bash
make start-db
```

## Run (in your local environment)

```bash
make run-local
```
