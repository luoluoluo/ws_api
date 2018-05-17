IMAGE_NAME=ws_api
# 编译
install:
	glide update
	go build --ldflags '-linkmode external -extldflags "-static"' -o ./main
	docker build -t ${IMAGE_NAME} .
	rm ./main
	docker save ${IMAGE_NAME} > ./dist/${IMAGE_NAME}
	cp ./Makefile ./dist/Makefile
	cp ./env.prod ./dist/env.prod
# 开发环境
rundev:
	docker run  -p 8000:8000 --env-file ./env.dev ${IMAGE_NAME}
# 正式环境
run:
	docker load < ./${IMAGE_NAME}
	docker run  -d -p 8000:8000 --env-file ./env.prod ${IMAGE_NAME} 1>>./${IMAGE_NAME}.log 2>>./${IMAGE_NAME}.log
