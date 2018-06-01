# echoudp

A command that echoes UDP payloads to the stdout.
It is specially designed to be used as a docker-compose service
to debug services that send UDP messages.
You will find Docker images for this project in [DockerHub](https://hub.docker.com/r/alcortesm/echoudp).

A PORT environment variable is required, with the port to listen to.
The process can be gracefully stopped with SIGTERM or SIGINT.

This repository includes an additional `healthcheck` command that can be used
to monitor the health of `echoudp` in docker or docker-compose.

## Installation

```
bash; go get github.com/alcortesm/echoudp/cmd/echoudp
```

## Example of use

```
bash; PORT=12345 echoudp
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

## Example of use (docker)
```
bash; docker run --env PORT=12345 -p'0.0.0.0:12345:12345/udp' alcortesm/echoudp:0.0.2
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

## Example of use (docker-compose)

```
bash; cat docker-compose.yml
version: '2.1'

services:
  echoudp:
    build: .
    environment:
      - PORT=12345
    healthcheck:
      test: ["CMD", "/bin/healthcheck"]
      interval: 250ms
      timeout: 250ms
      retries: 7

  client-example:
    image: subfuzion/netcat
    depends_on:
      echoudp:
        condition: service_healthy
    entrypoint: "/bin/ash"
    command: "-c 'echo -n hi! | nc -u -w1 echoudp 12345'"
bash;
bash;
bash; docker-compose build && docker-compose up
[...]
Successfully tagged echoudp_echoudp:latest
client-example uses an image, skipping
Starting echoudp_echoudp_1 ... 
Starting echoudp_echoudp_1 ... done
Recreating echoudp_client-example_1 ... 
Recreating echoudp_client-example_1 ... done
Attaching to echoudp_echoudp_1, echoudp_client-example_1
echoudp_1         | 2018-06-01T14:17:46.408Z echoudp started
echoudp_1         | 2018-06-01T14:17:46.408Z listening on 0.0.0.0:12345
echoudp_1         | 2018-06-01T14:17:47.307Z received: "hi!"
echoudp_client-example_1 exited with code 0
^CGracefully stopping... (press Ctrl+C again to force)
Stopping echoudp_echoudp_1 ... done
```

