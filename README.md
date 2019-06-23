# wait-for

Docker utility to wait for a database container to be available.

### Use

You can use wait-for binary for wait a MySQL database:

```sh
./wait-for --dbtype mysql \
           --host localhost \
           --port 3306 \
           --user root \
           --password mysql \
           --database mysql \
           --maxAttempts 100
```

And a PostgreSQL database:

```sh
./wait-for --dbtype postgres \
           --host localhost \
           --port 5432 \
           --user test \
           --database test \
           --password test \
           --maxAttempts 100
```

### Parameters

**--dbtype**: Set the database type. The default value is "postgres".

**--host**: Set the target database host. The default value is "localhost".

**--port**: Set the database port. The default value is "5432".

**--user**: Set the database user. The default value is "postgres".

**--password**: Set the database password. The default value is "postgres".

**--database**: Se the target database name. The default value is "postgres".

**--seconds**: Set the amount of seconds to wait for the database. The default value is "10".

**--maxAttempts**: Set the max attempts quantity. The default value is "3".

## Licensing

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0.html)
