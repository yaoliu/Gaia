

Gaia/[英文](README.md)
-------
Gaia 一个简单的网关服务

## Features
* 流量控制
* 熔断
* 负载均衡
* 路由
* 访问控制
* 防重访
* 认证（JWT,第三方认证)
* URL转换
* 


## 相关概念 
* API: 包含一组相关接口
* Cluster: 集群 包含一组相关服务
* Service: 服务 包含一组相关API
* Backend: 由一组后端节点组成
* Node: 提供服务的节点IP组成


##默认配置

```
global:
  logLevel: info
  clusterName: gaia
  clusterTag: gateway
  healthAddr: 9500
  healthPath: "/health"
```

###访问控制(filter/access)

```
filters:
  - name: "accessControl"
    services:
     - name: '{serviceName}'
       type: 'black'
       appIds: [110,123,114]
       ips: ['192.168.10.1']
```






###认证(filter/auth)

```
filters:
  - name: "auth"
```

###认证(filter/urlMapping/{clusterName}/{serviceName})

```
{
    "mappings":[
        {
            "url":"/aaaa/v1/bbbbb",
            "targetUrl":"/aaaa/bbbbb"
        }
    ]
}
```

###后端服务注册(backend/{clusterName}/{serviceName})

```
{
    "mappings":[
        {
            "url":"/aaaa/v1/bbbbb",
            "targetUrl":"/aaaa/bbbbb"
        }
    ]
}
```
