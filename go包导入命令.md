###by 十三小白

####**go 格式化时间** **2006-01-02 15:04:05 -0700 MST**

<**go环境设置**>

go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.cn,direct

<**gorm**>

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

<**mysql**>

go get github.com/gin-contrib-cors

<**mongodb**>

go get go.mongodb.org/mongo-driver/mongo

<**redis**>

go get github.com/go-redis/redis/v9

<**gin**>

go get github.com/gin-gonic/gin

<**gin_cookie session**>

go get github.com/gin-contrib/sessions

go get github.com/gin-contrib/sessions/cookie

<**jwt 跨域鉴权**>

go get github.com/dgrijalva/jwt-go

<**viper  配置管理**>

go get github.com/spf13/viper

<**cors 跨域访问 使用**>

go get github.com/gin-contrib-cors

<**logus 日志**>

go get github.com/sirupsen/logrus

go get github.com/lestrrat-go/file-rotatelogs

go get github.com/rifflock/lfshook

<**websocket**>

go get github.com/gorilla/websocket

<**captcha  验证码**>

go get github.com/dchest/captcha

<**e-mail 使用**>

go get github.com/jordan-wright/email

<**casbin 访问权限使用**>
go get github.com/casbin/casbin/v3
go get github.com/casbin/gorm-adapter/v3

<**carbon 时间处理**>

go get -u github.com/golang-module/carbon/v2

<**air 使用**>
go install github.com/cosmtrek/air@latest 

<**swaager 导入使用**>

 go get -u github.com/swaggo/swag/cmd/swag

 go get -u github.com/swaggo/gin-swagger 

 go get -u github.com/swaggo/gin-swagger/swaggerFiles

swag init 生成文档

------------------------------------------------------------------------------

<**protocol buffer (protobuf)**>

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

或者

go get github.com/golang/protobuf/protoc-gen-go

<**go-micro**>

go get -u github.com/micro/micro

go get -u github.com/micro/go-micro

go get -u github.com/micro/protoc-gen-micro

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest

go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

<**kafka 导入**>
go get github.com/shopify/sarama

<**常用HttpStatus状态**>：
HttpStatus.OK = 200;  
HttpStatus.BADREQUEST = 400;  
HttpStatus.FORBIDDEN = 403;  
HttpStatus.NOTFOUND = 404;  
HttpStatus.TIMEOUT = 408;  
HttpStatus.SERVERERROR = 500;  
————————————————----------------------------------------------------------

```
HttpStatus = {  
    //Informational 1xx  信息
    '100' : 'Continue',  //继续
    '101' : 'Switching Protocols',  //交换协议
    //Successful 2xx  成功
    '200' : 'OK',  //OK
    '201' : 'Created',  //创建
    '202' : 'Accepted',  //已接受
    '203' : 'Non-Authoritative Information',  //非权威信息
    '204' : 'No Content',  //没有内容
    '205' : 'Reset Content',  //重置内容
    '206' : 'Partial Content',  //部分内容

    //Redirection 3xx  重定向
    '300' : 'Multiple Choices',  //多种选择
    '301' : 'Moved Permanently',  //永久移动
    '302' : 'Found',  //找到
    '303' : 'See Other',  //参见其他
    '304' : 'Not Modified',  //未修改
    '305' : 'Use Proxy',  //使用代理
    '306' : 'Unused',  //未使用
    '307' : 'Temporary Redirect',  //暂时重定向

    //Client Error 4xx  客户端错误
    '400' : 'Bad Request',  //错误的请求
    '401' : 'Unauthorized',  //未经授权
    '402' : 'Payment Required',  //付费请求
    '403' : 'Forbidden',  //禁止
    '404' : 'Not Found',  //没有找到
    '405' : 'Method Not Allowed',  //方法不允许
    '406' : 'Not Acceptable',  //不可接受
    '407' : 'Proxy Authentication Required',  //需要代理身份验证
    '408' : 'Request Timeout',  //请求超时
    '409' : 'Conflict',  //指令冲突
    '410' : 'Gone',  //文档永久地离开了指定的位置
    '411' : 'Length Required',  //需要Content-Length头请求
    '412' : 'Precondition Failed',  //前提条件失败
    '413' : 'Request Entity Too Large',  //请求实体太大
    '414' : 'Request-URI Too Long',  //请求URI太长
    '415' : 'Unsupported Media Type',  //不支持的媒体类型
    '416' : 'Requested Range Not Satisfiable',  //请求的范围不可满足
    '417' : 'Expectation Failed',  //期望失败

    //Server Error 5xx  服务器错误
    '500' : 'Internal Server Error',  //内部服务器错误
    '501' : 'Not Implemented',  //未实现
    '502' : 'Bad Gateway',  //错误的网关
    '503' : 'Service Unavailable',  //服务不可用
    '504' : 'Gateway Timeout',  //网关超时
    '505' : 'HTTP Version Not Supported'  //HTTP版本不支持

};  
```

-------------------------------------------------------------------------------

<**Vue**>

npm install vue

<**Vue Router**>

npm install vue-router@4

<**vue_cli**>

npm install -g @vue/cli

<**axios**>

npm insrall axios

<**vue下载依赖**>

npm install --legacy-peer-deps

----------------------------------------------------------------------------------------

<**cmd中文乱码解决**>

CHCP 65001
