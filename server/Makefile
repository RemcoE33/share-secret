NAME="ghco-sercret"
PROJECT="ghco-invest"

buildui:
	yarn --cwd ./../ui/ build

run: buildui
	go run main.go

setproject:
	gcloud config set project ${PROJECT}

deploy: setproject buildui
	docker buildx build --platform linux/amd64 -t ${NAME} .
	docker tag ${NAME}:latest gcr.io/${PROJECT}/${NAME}
	docker push gcr.io/${PROJECT}/${NAME}
