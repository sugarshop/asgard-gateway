# asgard-gateway
A general-purpose API gateway for microservices to provide an HTTP endpoint, named Asgard.

The project uses [gin](https://github.com/gin-gonic/gin) as the `web` framework for the public access gateway layer.

Go RPC framework [cloudwego/kitex](https://github.com/cloudwego/kitex)

# Design Doc

- [Wiki Page](https://gamma.app/public/ChattyAI-l79uftz5bxwbdd8?mode=doc)
- [GPT API Doc](https://renaissancelabs101.notion.site/API-Access-ea86d8bd0e1345799db00bef03a92151?pvs=4)
- [Subscription Design Doc](https://renaissancelabs101.notion.site/Subscription-0ea7aa61c2514dafac72ea1764766fd0?pvs=4)
- [Comprehensive Cost Analyze Doc](https://renaissancelabs101.notion.site/Comprehensive-Cost-Analyze-of-Chatty-AI-System-b417ebe28a8542fe8bec4ee6f90438bb?pvs=4)

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
