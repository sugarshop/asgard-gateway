#!/bin/bash

MODULE_NAME="github.com/sugarshop/asgard-gateway"
BASE_IDL_DIR="../rpc/idl"
rm -rf ./kitex_gen

# Q: why `-module`
# A: https://github.com/cloudwego/kitex/blob/develop/docs/guide/basic-features/code_generation_cn.md#-module-module_name
kitex -module ${MODULE_NAME} -v ${BASE_IDL_DIR}/base/base.thrift
