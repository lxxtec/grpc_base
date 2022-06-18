# grpc-base

**gRPC 介绍**

gRPC是一个高性能开源通用rpc框架，面向移动和`HTTP/2`设计，支持多种语言

基于http2标准设计，带来诸如双向流，流控，头部压缩，单tcp连接上多服用请求等特性

关键特性：超时，重试，拦截器，命名解析，负载均衡，安全连接

传输：`http2+protobuf `

http1.0：提供长连接，请求回应的模式

http1.1：pipeline，可以发送多个请求得到多个回应

http2.0：`stream传输`



序列化： json xml msgPack `protobuf `

序列化要求：解压缩要快，数据流

json和xml 纯字符串，冗余信息太多了，解压缩较慢，msgPack会将k,v都压缩，而protobuf只压缩v，不压缩k，两端按数据流协议解析 



**gRPC service API**

1. unary api 一元普通模式，请求回应
2. client stream api 客户端流模式，可同时发送多个请求，服务端会汇总请求然后回应
3. server stream api 服务端流模式，客户端一个请求，服务端返回多个回应
4. biidirectional stream api 双端流，客户端多个请求，服务端返回多个回应
