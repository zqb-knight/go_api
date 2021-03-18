## 写在前面

这是一个学习go语言的练手项目，尽量将学到的知识进行运用。

- 使用go语言搭建一个HTTP的API
- 使用gin框架
- 使用redis, mysql进行数据存储
- 使用viper进行配置加载

## 前期准备
- 配置go环境

  文档：https://golang.org/doc/install

- docker安装redis, mysql镜像并启动容器
  - redis: `docker pull redis`
  - mysql: `docker pull mysql:5.7`

- 使用go mod配置依赖
  - gin: `go get github.com/gin-gonic/gin`
  - go-redis: `go get github.com/go-redis/redis`
  - viper: `go get github.com/spf13/viper`

## 进展
- [ ] 使用viper读取配置文件实现配置管理
- [ ] 编写全局错误码包
- [ ] 使用gin, mysql实现标签的RESFul API
- [ ] 实现JWT权限鉴定
- [ ] 为API增加redis缓存
- [ ] 使用Swqgger生成API文档
- [ ] 实现表单验证

## 技术点
1. 如何解决redis, mysql双写不一致问题
2. redis缓存雪崩，击穿，穿透的原理和解决办法
3. jwt的优点

