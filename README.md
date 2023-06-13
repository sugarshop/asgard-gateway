# asgard-gateway
A general-purpose API gateway for microservices to provide an HTTP endpoint, named Asgard.

The project uses [gin](https://github.com/gin-gonic/gin) as the `web` framework for the public access gateway layer.

## Layers

The project is divided into the following layers:
1. Handler
   Exposes HTTP endpoints to handle incoming requests and validate and assemble parameters without processing specific business logic.
2. Service
   Business logic layer.
3. Remote
   `RPC` client call package.

### Using the Dockerfile

Typically, you can follow these steps to use it:
1. In the image, you can enter `asgard-gateway`, then:
    1. Run `sh kitex.sh` to generate `kitex_gen`
    2. Run `pre-commit install` (execute once) -> coding -> `git commit -m 'tinyfix'`...
