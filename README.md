这是一个基于 [Go](https://go.dev/) 语言开发的项目，下面的运行代码将全部基于 [Docker](https://www.docker.com/) 运行演示

## 服务说明

- ADMIN：后台管理接口 
- WEB：官网请求交互接口

## ADMIN 接口

编译镜像

```bash
docker build -t noos/admin:1.0.0 -f docker/admin.Dockerfile .
```

运行容器

```bash
docker run \
  -d \
  --name noos-admin \
  --net uper \
  --ip 172.19.0.115 \
  -v $PWD/conf:/app/conf \
  noos/admin:1.0.0
```

## WEB 接口

编译镜像

```bash
docker build -t noos/web:1.0.0 -f docker/web.Dockerfile .
```

运行容器

```bash
docker run \
  -d \
  --name noos-web \
  --net uper \
  --ip 172.19.0.113 \
  -v $PWD/conf:/app/conf \
  noos/web:1.0.0
```
