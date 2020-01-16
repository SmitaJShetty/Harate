GO=go
GOBUILD= go build
APPNAME=ratelimiter
APP_BUILD_FOLDER=build
DOCKER_TAG=1.0.0

.PHONEY: build 
bld: 
	cd cmd/ && $(GOBUILD) -o ../$(APP_BUILD_FOLDER)/$(APPNAME)

build-linux: clean ## Prepare a build for a linux environment
	CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -o $(APP_BUILD_FOLDER)
	./$(APPNAME)

clean: 
	rm $(APP_BUILD_FOLDER)

docker-bld:
	docker build . -t $(APPNAME):$(DOCKER_TAG)
	
docker-run: docker-bld
	docker run -d -p 8080:8080 $(APPNAME):$(DOCKER_TAG)

run: bld
	cd $(APP_BUILD_FOLDER) && ./$(APPNAME)

test: bld
	go test ./...

