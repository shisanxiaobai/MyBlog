# myBlog

### 功能实现

外部分为前台后台两部分。
前台实现了主页、文章、动态及项目四项模块的展示。留言板，个人信息，播放音乐，评论，往日文章，分页，按分类检索，背景切换等功能。
后台实现了管理员登录，密码重置，跨域Token，对前台模块的增加，删除，编辑，列表，按条件查找操作。对评论，分类的管理。
内部实现了配置文件的实时监听，实时项目运行日志，实时开发日志，并做了持久化，分割及软连接处理。针对每一个接口实现swagger实时文档以及postman/airpost团队文档，对所有逻辑架构做了人性化，详细的注释说明。实现了错误处理、跨域、鉴权、热部署，Docker部署理等功能。

### 技术应用

前端：vue3框架，element-ui、vuesax界面组件库，tinymce富文本组件库等。
后端：mysql、gin、gorm框架，viper、ini配置管理，logurs、rotatelogs日志管理，swaggo接口文档，cors跨域，jwt跨域鉴权，air热部署，docker部署等

### 安装教程

#### docker下的安装（无需go环境，需要安装docker，并在docker中拉取启动mysql:8.0.28:并创建数据库myblog）：

修改config下的config.ini文件
进入文件所在的文件夹，执行docker build -f Dockerfile -t myblog .
执行 docker run -p 9000:9000 -d myblog
前端需要安装node.js、npm、vue、vue-cli以及vue Router，插件库需要下载依赖npm install --legacy-peer-deps
npm run serve运行后
前台网页访问localhost:8080，后台网页访问localhost:8081/admin  (端口顺序跟启动顺序相关)

#### 直接安装访问（需要配置go环境，并安装mysql：8.0.28，创建数据库myblog）：

修改config.ini文件
执行命令go mod tidy下载go依赖包
go run main.go执行
前端需要安装node.js、npm、vue、vue-cli以及vue Router，插件库需要下载依赖npm install --legacy-peer-deps
npm run serve运行后
前台网页访问localhost:8080，后台网页访问localhost:8081/admin  (端口顺序跟启动顺序相关)

### 使用说明

需要在windows下或者docker下安装mysql
如果主机名不是localhost，端口号想修改，需要修改config/config.ini并修改前端axios的URL后重新打包

swagger导入依赖后，swag init初始化 进入路由界面即可http://localhost:9000/swagger/index.html
授权token格式 Bearer+空格+登录返回reponse数据中token中的字符串

### 参与贡献

Fork 本仓库
新建 Feat_xxx 分支
提交代码
新建 Pull Request

### 项目实例

![前台](https://github.com/shisanxiaobai/MyBlog/blob/main/image/Snipaste_2022-11-11_22-48-58.png)
![后台](https://github.com/shisanxiaobai/MyBlog/blob/main/image/Snipaste_2022-11-11_22-49-56.png)
![swagger](https://github.com/shisanxiaobai/MyBlog/blob/main/image/Snipaste_2022-11-11_22-51-37.png)
