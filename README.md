# 该代码仓库使用go-graphql框架来构建web服务

## 1.层级调用结构及介绍

api层 -> domain层 -> infra层
各层级添加防腐层来进行参数限制

api: 接入graphql的入参和出参 进行参数校验 调用下层方法

auth: 做接入层的鉴权逻辑

common: 错误码声明、错误描述声明

cmd: 离线脚本 可直接调用domain层和infra层

docs: 整个仓库的文档存储 包含表结构

domain: 逻辑薄，仅做业务大结构体的构建，数据流转

infra: 初始化db、rpc连接方式，数据存储结构体建设

router: 配置路由信息

## 2.测试使用方式

curl --location --request POST 'http://localhost:8080/graphql' \
--header 'Content-Type: application/graphql' \
--header 'Cookie: pony' \
--data-raw 'query query {
    getUserInfo(userID: "pony"){
        userID, mobile, name, fcIDs
    }
}'