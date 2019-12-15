FROM gliderlabs/alpine:3.4

WORKDIR /flyway

ENV FLYWAY_VERSION=6.1.1 \
    WAIT_FOR_VERSION=v0.0.3 \
    MAVEN_REPO=http://central.maven.org/maven2

RUN apk upgrade --update

RUN apk add --no-cache bash

RUN apk-install openjdk8-jre ca-certificates wget \
  && wget "$MAVEN_REPO/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}.tar.gz" \
  && tar -xzf flyway-commandline-${FLYWAY_VERSION}.tar.gz \
  && mv flyway-${FLYWAY_VERSION}/* . \
  && rm flyway-commandline-${FLYWAY_VERSION}.tar.gz \
  && ln -s /flyway/flyway /usr/local/bin/flyway

# Adding wait-for
RUN wget "https://github.com/ArcanjoQueiroz/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for" \
  && chmod u+x wait-for && mv wait-for /usr/local/bin/wait-for \
  && echo -ne "- with Flyway $FLYWAY_VERSION\n" >> /root/.built

CMD ["flyway", "--help"]