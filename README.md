# asgard-gateway
An API gateway for microservices to provide HTTP endpoint, named Asgard.

公网接入的通用网关层，项目使用[gin](https://github.com/gin-gonic/gin)作为`web`框架。

## Layers

项目分如下几层：
1. handler
    对外暴露HTTP endpoint，负责处理接收请求，处理上行参数校验、拼装，不处理具体业务
2. service
    业务层
3. remote
    `RPC`调用client包

### Use the Dockerfile

通常，你可按如下步骤来使用：
1. 在镜像里，你可以进入`asgard-gateway`， 然后:
    1. 执行`sh kitex.sh`来生成`kitex_gen`
    2. 执行`sh genmock.sh`来生成mock代码
    3. `pre-commit install`(执行一次) -> coding -> `git commit -m 'tinyfix'` ...
