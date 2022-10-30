# myBlog

### 功能实现
外部分为前台后台两部分。
前台实现了主页、文章、动态及项目四项模块的展示。留言板，个人信息，播放音乐，评论，往日文章，分页，按分类检索，背景切换等功能。
后台实现了管理员登录，密码重置，跨域Token，对前台模块的增加，删除，编辑，列表，按条件查找操作。对评论，分类的管理。
内部实现了配置文件的实时监听，实时项目运行日志，实时开发日志，并做了持久化，分割及软连接处理。针对每一个接口实现swagger生成，也通过了postman/airpost测试，对所有逻辑架构做了人性化，详细的注释说明。实现了错误处理、跨域、鉴权、热部署，Docker部署理等功能。

### 技术应用
前端：vue3框架，element-ui、vuesax界面组件库，tinymce富文本组件库等。
后端：gin、gorm框架，viper、ini配置管理，logurs、rotatelogs日志管理，swaggo接口文档，cors跨域，jwt跨域鉴权，air热部署，docker部署等

### 安装教程
#### docker下的安装（无需go环境，需要安装docker，并在docker中拉取启动mysql:8.0.28:并创建数据库myblog）：

修改config下的config.ini文件
进入文件所在的文件夹，执行docker build -f Dockerfile -t myblog .
执行 docker run -p 8000:8000 -d myblog
前台网页访问localhost:8000，后台网页访问localhost:8000/admin


#### 直接安装访问（需要配置go环境，并安装mysql：8.0.28，创建数据库myblog）：

修改config.ini文件
执行命令go mod tidy下载go依赖包
go run main.go执行
前台网页访问localhost:8000，后台网页访问localhost:8000/admin

### 使用说明
需要在windows下或者docker下安装mysql
如果主机名不是localhost，需要修改config/config.ini并修改前端axios的baseURL后重新打包

### 参与贡献
Fork 本仓库
新建 Feat_xxx 分支
提交代码
新建 Pull Request
