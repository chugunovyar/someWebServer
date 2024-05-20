# someWebServer

# Description

The service is written to demonstrate programming skills in golang 1.21



# Environment variables

| Name          | Default Value | Description                                                        |
|---------------|--------------|--------------------------------------------------------------------|
|  POSTGRES_NAME             | postgres     | Database name                                         |
|       POSTGRES_USER        | postgres     | Database Username                                     |
|   POSTGRES_PASSWORD            | postgres     | Database password                                 |
|   POSTGRES_PORT            | 5432         | Database port                                         |
|  DB_HOST             | db\localhost | Hostname for database. Depends on how you run the server    |
|  LOG_LEVEL           | debug| logging level can be set one of trace, debug, info, error, warning |


# Running
Can be launched via Docker:
```shell
docker compose up --build
```

or directly, but you must set env vars.
```shell
go run .
```

