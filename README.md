# wait-for

_Wait-for_ is a binary used to wait for a resource (like a _database_ or a _message broker_) to be available. It supports _PostgreSQL_, _MySQL_, _MongoDB_ databases and _RabbitMQ_ message broker.

## Installing

Here is an example of _wait-for_ installation in a _Dockerfile_:

```dockerfile
ENV WAIT_FOR_VERSION=v0.0.4
RUN wget "https://github.com/arcanjoaq/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for" && chmod u+x wait-for
  ```

## Usage

Basically, the option **--type** determines the target resource. Here are some examples of _wait-for_:

### MySQL

```sh
./wait-for --type mysql \
           --host localhost \
           --port 3306 \
           --user root \
           --password mysql \
           --name mysql \
           --maxAttempts 100
```

### PostgreSQL

```sh
./wait-for --type postgres \
           --host localhost \
           --port 5432 \
           --user test \
           --password test \
           --name test \
           --maxAttempts 100
```

### MongoDB

```sh
./wait-for --type mongodb \
           --host localhost \
           --port 27017 \
           --user test \
           --password test \
           --name test \
           --maxAttempts 100
```

### RabbitMQ

```sh
./wait-for --type rabbitmq \
           --host localhost \
           --port 5672 \
           --user test \
           --password test \
           --name '/' \
           --maxAttempts 100
```

## Options

**--type**: Set _resource type_.

**--host**: Set _target host_. The default value is "localhost".

**--port**: Set _resource port_. The default value depends on resource type.

**--user**: Set _resource user_.

**--password**: Set _resource password_.

**--name**: Se _resource name_. It is _database name_ or _virtual host name_.

**--seconds**: Set _number of seconds_ to wait for a resource. The default value is "10".

**--maxAttempts**: Set _max attempts quantity_. The default value is "3".

## Linting

```sh
make lint
```

## Building

```sh
make build
```

## Licensing

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0.html)
