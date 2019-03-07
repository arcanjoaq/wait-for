# wait-for

Docker utility to wait for a container to be available.

## wait-for-postgres

### Use

You can use wait-for-postgres this way:

```sh
$ ./wait-for-postgres
```

or in a more completed form:

```sh
$ ./wait-for-postgres --host localhost --port 5432 --user postgres --password postgres --database postgres --seconds 10 --maxAttempts 3
```

### Parameters

**--host**: Set the target database host. The default value is "localhost".

**--port**: Set the database port. The default value is "5432".

**--user**: Set the database user. The default value is "postgres".

**--password**: Set the database password. The default value is "postgres".

**--database**: Se the target database name. The default value is "postgres".

**--seconds**: Set the amount of seconds to wait for the database. The default value is "10".

**--maxAttempts**: Set the max attempts quantity. The default value is "3".


