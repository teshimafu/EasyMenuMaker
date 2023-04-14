
PROJECT_NAME=lazyPM
DOCKER_PROJECT_NAME=lazy_pm

# ビルドコマンドを定義
build:
	go build -o $(PROJECT_NAME)

# cleanコマンドを定義
clean:
	rm -f $(PROJECT_NAME)

run: build
	./$(PROJECT_NAME)

docker-build:
	docker build -t $(DOCKER_PROJECT_NAME) .

docker-run: docker-build
	docker run -it --rm -p 8080:8080 $(DOCKER_PROJECT_NAME)