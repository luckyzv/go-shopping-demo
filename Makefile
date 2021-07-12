OUT="out"

build:
	@go build -o ${OUT}

run:
	@go run ./main.go

swag:
	swag init

docker-login:
	sh ./docker/docker-login.sh

docker-build: docker-login
	@docker build -t registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version} .

docker-push: docker-build
	@docker push registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version}

deploy-staging:
	cd ./docker && docker-compose -f ./redis/docker-compose.yml -f ./sentinel/docker-compose.yml -f ./docker-compose.staging.yml up -d

deploy-prod:
	cd ./docker && docker-compose -f ./redis/docker-compose.yml -f ./sentinel/docker-compose.yml -f ./docker-compose.prod.yml up -d
