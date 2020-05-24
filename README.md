# GO notes

### Development environment

Set up environment. App will be available on `http://localhost:8080`.

```sh
make up
```

Stop local environment

```sh
make up
```

### Tests

All tests are ran in container

```sh
make test
```

### Migrations

Run all UP migrations. Local environment has to be on. 

```sh
make migrate-db-up
```

Rollback last migration

```sh
make migrate-db-down
```