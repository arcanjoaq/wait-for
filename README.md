# wait-for

This is a Docker utility to wait for a database container to be available. Currently the wait-for binary supports **PostgreSQL 11** and **MySQL 8.0.16** databases.

### Installing

Here is an example of wait-for installation in a Dockerfile:

```dockerfile
ENV WAIT_FOR_VERSION=v0.0.2
RUN wget "https://github.com/ArcanjoQueiroz/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for" && chmod u+x wait-for
  ```


### Usage

Basically, the option **--dbtype** determines the target database. Here are some examples of wait-for:

#### MySQL

```sh
./wait-for --dbtype mysql \
           --host localhost \
           --port 3306 \
           --user root \
           --password mysql \
           --database mysql \
           --maxAttempts 100
```

#### PostgreSQL

```sh
./wait-for --dbtype postgres \
           --host localhost \
           --port 5432 \
           --user test \
           --database test \
           --password test \
           --maxAttempts 100
```

### Options

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
