OUT="out"

build:
	@go build -o ${OUT}

run:
	@go run ./

login:
	sh ./docker-login.sh

docker-build: login
	@docker build -t registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version} .


docker-push: docker-build
	@docker push registry.cn-shanghai.aliyuncs.com/luckyziv/go-shopping-demo:${Version}
