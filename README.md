# echo-udp

A command that echoes UDP payloads to the stdout.

It is specially designed to be used as a docker-compose service
to debug services that send UDP messages.

A PORT environment variable is required, with the port to listen to.
It includes an internal healthcheck
and can be gracefully stopped with SIGTERM or SIGINT.

## Installation

```
bash; go get github.com/alcortesm/echo-udp/cmd/echo-udp
```

## Example of use

```
bash; PORT=12345 echo-udp
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

## Example of use (docker)
```
bash; docker run --env PORT=12345 -p'0.0.0.0:12345:12345/udp' alcortesm/echo-udp:0.0.1
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

## Example of use (docker-compose)

```
version: '2.1'

services:
  sut:
    build: .
    environment:
      - PORT=12345

  tester:
    image: subfuzion/netcat
    depends_on:
      sut:
        condition: service_healthy
    entrypoint: "/bin/ash"
    command: "-c 'nc -u -w1 sut 12345 < /etc/alpine-release'"
```

