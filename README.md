# wait-for

[![Build Status](https://travis-ci.com/arcanjoaq/wait-for.svg?branch=master)](https://travis-ci.com/arcanjoaq/wait-for)

This is a Docker utility to wait for a database container to be available. Currently the wait-for binary supports **PostgreSQL 11**, **MySQL 8.0.16**, **MongoDB 3.4.19+** databases and **RabbitMQ 3.7+** message broker.

### Installing

Here is an example of wait-for installation in a Dockerfile:

```dockerfile
ENV WAIT_FOR_VERSION=v0.0.3
RUN wget "https://github.com/ArcanjoQueiroz/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for" && chmod u+x wait-for
  ```


### Usage

Basically, the option **--type** determines the target resource. Here are some examples of wait-for:

#### MySQL

```sh
./wait-for --type mysql \
           --host localhost \
           --port 3306 \
           --user root \
           --password mysql \
           --name mysql \
           --maxAttempts 100
```

#### PostgreSQL

```sh
./wait-for --type postgres \
           --host localhost \
           --port 5432 \
           --user test \
           --password test \
           --name test \
           --maxAttempts 100
```

#### MongoDB

```sh
./wait-for --type mongodb \
           --host localhost \
           --port 27017 \
           --user test \
           --password test \
           --name test \
           --maxAttempts 100
```

#### RabbitMQ

```sh
./wait-for --type rabbitmq \
           --host localhost \
           --port 5672 \
           --user test \
           --password test \
           --name '/' \
           --maxAttempts 100
```

### Options

**--type**: Set the resource type. The default value is "postgres".

**--host**: Set the target host. The default value is "localhost".

**--port**: Set the resource port. The default value is "5432".

**--user**: Set the resource user. The default value is "postgres".

**--password**: Set the resource password. The default value is "postgres".

**--name**: Se the resource name. It is the database name or virtual host name. The default value is "postgres".

**--seconds**: Set the amount of seconds to wait for resource. The default value is "10".

**--maxAttempts**: Set the max attempts quantity. The default value is "3".

## Licensing

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0.html)
