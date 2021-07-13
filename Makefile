OUT="out"

# 打包项目
build:
	@go build -o ${OUT}

# 运行项目
run:
	@go run ./main.go

# 生成swagger文档
swag:
	swag init

# 登录阿里云docker仓库
docker-login:
	sh ./docker/docker-login.sh

# 创建docker镜像
docker-build: docker-login
	@docker build -t registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version} .

# 推送docker镜像到阿里云docker仓库
docker-push: docker-build
	@docker push registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version}

# docker-compose构建redis主从
docker-build-redis:
	cd ./docker && docker-compose -f ./redis/docker-compose.yml up -d

# docker-compose构建redis哨兵
docker-build-sentinel: docker-build-redis
	cd ./docker && docker-compose -f ./sentinel/docker-compose.yml up -d

# 部署整个项目到staging
deploy-staging: docker-build-sentinel
	cd ./docker && docker-compose -f ./docker-compose.staging.yml up -d

# 部署整个项目到prod
deploy-prod: docker-build-sentinel
	cd ./docker && docker-compose -f ./docker-compose.prod.yml up -d
