# Notification pusher

Very simple service that send message using Firebase.

## Why Golang?
Just to try this language. 

Whenever you would take Golang as your app development language you have to take many things into account. And do not forget
about your coworkers that do not know Golang at all.

Some considerations - why to use Golang for this service:
1) new and interesting language
2) goroutine model to make your service very effective on high load (especially because of high IO load in this service)
3) your coworkers write on Golang :P

## How to start
You have to specify configuration in file or using environment variables

```yaml
server:
  run_mode: # what Gin mode you will use - release, debug, test as possible variants. release is for production; env var - SERVER_RUN_MODE
  port: # at which port we will bind our server; env var - SERVER_PORT
  host: # at what host we will bind our server - localhost is number one for your decisions; env var - SERVER_HOST
  read_timeout: 3s #  ReadTimeout is the maximum duration for reading the entire request, including the body.
  write_timeout: 3s # WriteTimeout is the maximum duration before timing out writes of the response. It is reset whenever a new request's header is read.

database:
  connection_string: # full connection to postgres in PQ format; for example, postgres://test:test@localhost:5432/postgres?sslmode=disable; env var - DATABASE_CONN_STRING

service:
  firebase_credentials: # base64 encoded string with firebase credentials file for Admin SDK; more information in Firebase Admin SDK documentation; env var - FIREBASE_CREDENTIALS.
```

File is presented in work directory - `config.yml`.
After you have changed this parameters you can start service

```bash
go run .
```

## Tech stack

1) Gin as web framework
2) Firebase SDK to work with notifications
3) sqlc to generate DAO objects for SQL integration
4) atlas to make migrations of SQL tables

### sqlc support

We have file `sqlc.yaml`, that contains parameters to process DAO creation. Also we have `schema.sql` and `query.sql` files that describe all requests that are used in connection with SQL.

After we collect all information to generate DAO we can execute next command in directory with `sqlc.yaml`

```bash
sqlc generate
```

after that in models you will see generated DAO.

### atlas support

For atlas we use `migrations` directory. All migrations must be applied using separate CI/CD steps (service do not upgrade tables structure itself).

For now we have some migrations. If we need to add more migrations we can 
1) make changes to our dev database (test db)
2) execute next command on root directory of project

```bash
atlas migrate diff new_very_logic \           
  --dir file://migrations \
  --to postgres://test:test@localhost:5432/postgres \
  --dev-url postgres://test:test@localhost:5432/test
```

`to` parameter set current db version. And `dev-url` set target db version. After command execution we will have another migration file that can be executed.