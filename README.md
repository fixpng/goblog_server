# 项目介绍

基于vue3和gin框架开发的前后端分离个人博客系统，包含md格式的文本编辑展示，点赞评论收藏，新闻热点，匿名聊天室，文章搜索等功能。

---

## 在线预览

博客前台：[http://www.fixpng.top/](http://www.fixpng.top/)
博客后台：[http://www.fixpng.top/admin/](http://www.fixpng.top/admin/)
API文档： [http://www.fixpng.top/swagger/index.html](http://www.fixpng.top/swagger/index.html)
> 网站域名已通过ICP备案
> 后台需要登陆访问，可用测试用户：`test / test`，也可自行注册账号
> 页面暂未适配移动端

![在这里插入图片描述](https://img-blog.csdnimg.cn/ace1b6b4a5914582bea89b379f835d34.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/d98e8915af6846e486d5e315ae7b1677.png)




---

## 技术介绍

前端代码：[https://github.com/fixpng/gvb_web](https://github.com/fixpng/gvb_web)
后端代码：[https://github.com/fixpng/gvb_server](https://github.com/fixpng/gvb_server)
> 前端部分是现学现卖的，技术和内容仍有许多不足，后面有时间会慢慢优化完善。

#### 开发环境

|开发工具 |功能描述			
|--|--|
|GoLand |后端开发		
|PyCharm |前端开发			
|Postman |API测试			
|Docker Desktop |依赖软件运行
|MobaXterm |Linux 远程工具

|开发环境 |版本 |
|--|--|
|GoLang | 1.19 |
|npm | 9.6.2 |
|MySQL | 5.7 |
|Elasticsearch | 7.12.0 |
|Redis | 7.0.8 |

---

#### 技术栈

> 这里只写一些主流的通用技术，详细第三方库：前端参考 package.json 文件，后端参考 go.mod 文件

|功能描述|前端|官方地址|
|--|--|--|
|Vue3框架|vuejs|https://cn.vuejs.org/
|Vue组件|ant-design-vue|https://next.antdv.com/docs/vue/introduce-cn/
|Mark Down|md-editor-v3|https://imzbf.github.io/md-editor-v3/
|状态管理工具|pinia|https://pinia.vuejs.org/|
|构建工具|vite|https://cn.vitejs.dev/
|可视化图表库|echarts|https://echarts.apache.org/zh/index.html|

|功能描述|后端|官方地址|
|--|--|--|
|GO语言|golang|https://github.com/golang/go
|WEB框架|gin|https://gin-gonic.com/zh-cn/docs/
|API文档|swaggo|https://github.com/swaggo
|ORM 库|gorm|https://github.com/go-gorm/gorm
|日志库|logrus|https://github.com/sirupsen/logrus
---

#### 数据库表关系

![1](https://img-blog.csdnimg.cn/a1bd4b52f8864e2aa2cfcf29e279413a.png)

---

## 本地运行

> 自行安装 Golang、Node、MySQL、Redis 、Elasticsearch 环境
> Golang 安装参考官方文档
> Node 建议安装使用 `https://nodejs.org/zh-cn` 的长期维护版
> MySQL、Redis、Elasticsearch 建议使用 Docker 运行

后端项目运行：

```powershell
# 1、启动MySQL、Redis、Elasticsearch,其中mysql需要新建一个库
# 2、修改项目运行的配置文件 settings.yaml

# 3、初始化运行环境
go mod tidy 				# 下载当前项目所依赖的包
go run main.go -db 			# mysql建表
go run main.go -es create	# elasticsearch建索引

# 4、mysql插入菜单表数据
INSERT INTO gvb.menu_models (id, created_at, updated_at, title, path, slogan, abstract, abstract_time, banner_time, sort) VALUES (1, NOW(), NOW(), '首页', '/', '众神眷恋的幻想乡', '天寒地冻路远马亡又何妨', 5, 5, 1);
INSERT INTO gvb.menu_models (id, created_at, updated_at, title, path, slogan, abstract, abstract_time, banner_time, sort) VALUES (2, NOW(), NOW(), '新闻', '/news', '新闻三分钟，知晓天下事', '震惊!男人看了会沉默，女人看了会流泪!不转不是中国人!',  5, 5, 2);
INSERT INTO gvb.menu_models (id, created_at, updated_at, title, path, slogan, abstract, abstract_time, banner_time, sort) VALUES (3, NOW(), NOW(), '文章搜索', '/search', '文章搜索', '文章搜索',  5, 5, 3);
INSERT INTO gvb.menu_models (id, created_at, updated_at, title, path, slogan, abstract, abstract_time, banner_time, sort) VALUES (4, NOW(), NOW(),'聊天室', '/chat_group', '聊天室', '聊天室',  5, 5, 4);

# 5、创建第一个用户，后续可在前端创建或注册
go run main.go -u admin		# 管理员
go run main.go -u user		# 普通用户

# 6、启动项目
go run main.go
```

前端项目运行：

```powershell
# 下载当前项目所依赖的包
npm i
# 启动项目
npm run dev
```

---

## 线上部署（Linux）| 待完善

本项目线上部署目录结构如下，必需的目录及文件：

```bash
gvb									# 项目启动必需的目录及文件
├── deploy							# 该项目依赖的数据库及中间件，均用docker启动
│   ├── docker-compose.yml			# Docker Compose 配置文件
│   ├── .env						# 环境配置文件
│   └── gvb
│       ├── elasticsearch
│       │   ├── config 
│		│	│	└── elasticsearch.yml # Elasticsearch 配置文件
│       │   ├── data
│       │   └── plugins
│       ├── mysql
│       ├── nginx
│       │   ├── conf
│       │   │	└── nginx.conf		# nginx配置文件
│       │   └── html
│       └── redis
├── gvb_server						# 后端代码
│   ├── docs
│   ├── uploads
│   └── main						# go打包文件
└── gvb_web							# 前端代码
    └── dist						# npm打包文件
```

---

#### 安装Docker

> 只需要提前安装好docker运行环境

1.安装 pip：

```Shell
yum -y install epel-release
yum -y install python-pip
 
#升级
pip install --upgrade pip
```

2.安装Docker-Compose：

```Shell
yum install python-devel

pip install --ignore-installed requests

pip install docker-compose
#检查是是否成功：
docker-compose -version
```

3.安装Docker:

```Shell
yum provides '*/applydeltarpm'  
yum install deltarpm -y

yum install docker-ce -y

curl -sSL https://get.daocloud.io/docker | sh
```

---

#### 依赖软件准备

> 只需要提前安装好docker运行环境
> 用 Docker 启动的项目依赖软件：mysql，redis，elasticsearch，nginx
> 可直接拷贝[后端代码](http://www.fixpng.top/admin/)中的deploy文件夹自行修改

docker-compose.yml 和 .env 文件放在部署服务器的 deploy 目录下
准备docker-compose.yml文件

```powershell
version: "1"

networks:
  gvb-network:
    driver: bridge
    ipam:
      config:
        - subnet: ${SUBNET}

services:
  gvb-redis:
    image: redis:7.0.8
    container_name: gvb-redis
    restart: always
    volumes:
      - ${GVB_DATA_DIRECTORY}/redis/data:/data
    ports:
      - ${REDIS_PORT}:6379 # 自定义的是暴露出去的端口, Redis 容器内运行固定为 6379
    command: redis-server --requirepass ${REDIS_PASSWORD} --appendonly yes
    networks:
      gvb-network:
        ipv4_address: ${REDIS_HOST}

  gvb-mysql:
    image: mysql:5.7
    container_name: gvb-mysql
    restart: always
    volumes:
      - ${GVB_DATA_DIRECTORY}/mysql/data:/var/lib/mysql
    environment:
      - MYSQL_ROOT_PASSWORD = ${MYSQL_ROOT_PASSWORD} # root 账号的密码
      - MYSQL_DATABASE = ${MYSQL_DATABASE} # root 账号的密码
      - MYSQL_USER: ${MYSQL_USER}
      - MYSQL_PASSWORD: ${MYSQL_PASSWORD}
      - TZ=Asia/Shanghai
    command:
      --max_connections=1000
      --character-set-server=utf8mb4
      --collation-server=utf8mb4_general_ci
    ports:
      - ${MYSQL_PORT}:3306 # 自定义的是暴露出去的端口, MySQL 容器内运行固定为 3306
    networks:
      gvb-network:
        ipv4_address: ${MYSQL_HOST}
        
  gvb-elasticsearch:
    image: elasticsearch:7.12.0
    container_name: gvb-elasticsearch
    restart: always
    volumes:
      - ${GVB_DATA_DIRECTORY}/elasticsearch/data:/usr/share/elasticsearch/data
      - ${GVB_DATA_DIRECTORY}/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml
      - ${GVB_DATA_DIRECTORY}/elasticsearch/plugins:/usr/share/elasticsearch/plugins
    environment:
      - discovery.type=single-node
      - ES_JAVA_OPTS=-Xms84m -Xmx512m
    ports:
      - ${ELASTICSEARCH_PORT01}:9200 # 自定义的是暴露出去的端口, elasticsearch 容器内运行固定为 9200和9300
      - ${ELASTICSEARCH_PORT02}:9300
    networks:
      gvb-network:
        ipv4_address: ${ELASTICSEARCH_HOST}
        
  gvb-nginx:
    image: nginx:latest
    container_name: gvb-nginx
    restart: always
    volumes:
      - ${GVB_DATA_DIRECTORY}/nginx/conf/nginx.conf:/etc/nginx/nginx.conf
      - ${GVB_DATA_DIRECTORY}/nginx/logs:/var/log/nginx
      - ${NGINX_GVB_WEB}:${NGINX_GVB_WEB} # 前端web目录
      - ${NGINX_GVB_SERVER}:${NGINX_GVB_SERVER} # 后端server目录
    ports:
      - ${NGINX_PORT}:80 # 自定义的是暴露出去的端口, nginx 容器内运行固定为 80
    networks:
      gvb-network:
        ipv4_address: ${NGINX_HOST}

```

准备 .env 文件（相关参数自行修改）

```bash
# https://docs.docker.com/compose/migrate/
# docker-compose.yml 同目录下的 .env 文件会被加载为其环境变量

# COMPOSE_PROJECT_NAME=gin-vue-blog

# 数据存储的文件夹位置 (默认在当前路径生成 gvb 文件夹)
GVB_DATA_DIRECTORY=./gvb

# Redis
REDIS_PORT=6379
REDIS_PASSWORD=12345QWERT

# MySQL
MYSQL_PORT=3306
MYSQL_ROOT_PASSWORD=12345QWERT
MYSQL_DATABASE=gvb_db
MYSQL_USER=gvb
MYSQL_PASSWORD=12345QWERT

# Elasticsearch
ELASTICSEARCH_PORT01=9200
ELASTICSEARCH_PORT02=9300

# Nginx
NGINX_PORT=80
NGINX_GVB_WEB=../gvb/gvb_web			# 给nginx映射的前端路径，和nginx.conf里路径对应
NGINX_GVB_SERVER=../gvb/gvb_server	# 给nginx映射的后端路径，和nginx.conf里路径对应

# Docker Network (一般不需要变, 除非发生冲突)
SUBNET=172.12.0.0/24
REDIS_HOST=172.12.0.2
MYSQL_HOST=172.12.0.3
ELASTICSEARCH_HOST=172.12.0.4
NGINX_HOST=172.12.0.5
```

Elasticsearch 配置文件：elasticsearch.yml

```powershell
http.host: 0.0.0.0
```

nginx配置文件：nginx.conf

```powershell

#user  root;
worker_processes  auto;

error_log  /var/log/nginx/error.log notice;
pid        /var/run/nginx.pid;

events {
	worker_connections  1024;
}

http {
	include       /etc/nginx/mime.types;
	default_type  application/octet-stream;

	log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
					'$status $body_bytes_sent "$http_referer" '
					'"$http_user_agent" "$http_x_forwarded_for"';

	access_log  /var/log/nginx/access.log  main;
	sendfile        on;
	#tcp_nopush     on;

	client_max_body_size 8M; #上传文件大小限制
	keepalive_timeout  65;

	server {
	listen 80;      # http
	#listen 443 ssl; # https
	server_name www.fixpng.top; # 域名

	##填写证书文件名称
	#ssl_certificate cert/server.pem;
	##填写证书私钥文件名称
	#ssl_certificate_key cert/server.key;
	#
    #ssl_session_cache shared:SSL:50m;
	#
	##自定义设置使用的TLS协议的类型以及加密套件（以下为配置示例，请您自行评估是否需要配置）
	##TLS协议版本越高，HTTPS通信的安全性越高，但是相较于低版本TLS协议，高版本TLS协议对浏览器的兼容性较差。
	#ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
    #ssl_protocols SSLv3 SSLv2 TLSv1 TLSv1.1 TLSv1.2;
	#
	##表示优先使用服务端加密套件。默认开启
	#ssl_prefer_server_ciphers on;

		location / {
			root   /www/wwwroot/gvb/gvb_web/dist/; #访问路径，相当于Tomcat的ROOT，这里自己配
			index  index.html index.htm; #访问index
			try_files $uri $uri/ /index.html; #解决刷新404问题
		}

		location /wsUrl/ {
			rewrite ^/wsUrl/(.*)$ /$1 break;   # 长连接时间
			proxy_pass http://172.12.0.1:8080/api/;
			proxy_http_version 1.1;
			proxy_set_header Upgrade $http_upgrade;
			proxy_set_header Connection "Upgrade";
			proxy_redirect off;
			proxy_set_header Host $host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Host $server_name;
			proxy_read_timeout 3600s;  # 长连接时间
		}

		location /api/ {
		# rewrite ^/(api/.*) /$1 break;
		proxy_set_header Host $host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header REMOTE-HOST $remote_addr;
		proxy_pass http://172.12.0.1:8080/api/;
		}

		location /uploads/ {
		# proxy_pass http://172.12.0.1:8080/uploads/;
		alias /www/wwwroot/gvb/gvb_server/uploads/;
		}

		location /swagger/ {
		proxy_pass http://172.12.0.1:8080/swagger/;
		}
    }
}
```

---

#### 应用程序准备

后端项目打包生成的main文件、docs文件夹、settings.yaml、uploads文件夹复制至部署服务器的gvb_server目录

```bash
# 生成api文档
swag init 

# 后端go打包
set GOARCH=amd64
set GOOS=linux
go build -o main 
```

前端项目打包生成的dist文件夹及其文件复制至部署服务器的gvb_web目录

```bash
# 前端npm打包
npm run build
```

---

#### 启动应用

修改好各项配置

```bash
# docker compose 启动依赖软件
cd xxxxx/gvb/deploy/
docker compose up -d

#启动后端应用
cd xxxxx/gvb/gvb_server/
nohup ./main &
```

访问应用

