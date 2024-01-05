SHELL=/bin/bash
GIT_HASH=`git rev-parse --short HEAD`
BUILD_DATE=`date +%FT%T%z`

BIN=bin/bankapp.app
SRC=./

LDFLAGS=-w -s -X main.GitHash=${GIT_HASH} -X main.BuildDate=${BUILD_DATE}

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
export COSIGN_PASSWORD=password

.PHONY: clean build

vet:
	go vet ./...

# lint: vet
# 	golangci-lint run

# build: vet
# 	mkdir -p bin/
# 	go build -ldflags "${LDFLAGS}" -o "${BIN}" "${SRC}"

docker-build: 
	docker build \
		--no-cache=true \
		--build-arg BUILD_REV=${GIT_HASH} \
		--build-arg BUILD_DATE=${BUILD_DATE} \
		-t bankapp:v1 \
		-f Dockerfile .

docker-tag: 
	docker tag bankapp:v1 moazrefat/bankapp:v1

docker-push: 
	docker push moazrefat/bankapp:v1

# cleanup: 
# 	for i in `docker images --format "{{.Repository}}:{{.Tag}}|{{.ID}}" |grep -i bankapp |cut -f2 -d'|'`; 
# 	do docker rmi -f $i; 
# 	done

docker: 
	docker build \
		--no-cache=true \
		--build-arg BUILD_REV=${GIT_HASH} \
		--build-arg BUILD_DATE=${BUILD_DATE} \
		-t bankapp:v1 \
		-f Dockerfile .
	docker tag bankapp:v1 moazrefat/bankapp:v1
	docker push moazrefat/bankapp:v1
	cosign sign --key ../../cosign.key moazrefat/bankapp:v1

docker-distroless: 
	docker build \
		--no-cache=true \
		--build-arg BUILD_REV=${GIT_HASH} \
		--build-arg BUILD_DATE=${BUILD_DATE} \
		-t bankapp:distroless \
		-f Dockerfile.distroless .
	docker tag bankapp:distroless moazrefat/bankapp:distroless
	docker push moazrefat/bankapp:distroless
	cosign sign --yes --key ../../cosign.key moazrefat/bankapp:distroless

docker-alpine: 
	docker build \
		--no-cache=true \
		--build-arg BUILD_REV=${GIT_HASH} \
		--build-arg BUILD_DATE=${BUILD_DATE} \
		-t bankapp:alpine \
		-f Dockerfile.alpine .
	docker tag bankapp:alpine moazrefat/bankapp:alpine
	docker push moazrefat/bankapp:alpine
	cosign sign --yes --key ../../cosign.key moazrefat/bankapp:alpine

clean:
	[ -f "${BIN}" ] && rm "${BIN}"