#################### 1
# FROM golang:1.18 AS builder

# WORKDIR /workspace
# COPY go.mod go.mod
# COPY go.sum go.sum

# RUN go mod download
# # RUN go mod tidy
# # RUN go mod verify

# COPY main.go main.go
# COPY assets/ assets/
# COPY pkg/ pkg/
# COPY views/ views/

# RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o ./bankapp main.go

# FROM gcr.io/distroless/static:nonroot

# WORKDIR /

# COPY --from=builder /workspace/bankapp .

# EXPOSE 8080

# ENTRYPOINT ["/bankapp", "-p", "8080"]


#################### 2
FROM golang:1.21.5

WORKDIR /
COPY . .

RUN GOOS=linux go build -gcflags "all=-N -l" -o /bankapp main.go

EXPOSE 8080

ENTRYPOINT ["/bankapp", "-p", "8080"]

#################### 3
# FROM golang:1.18 AS build-stage

# WORKDIR /go/src/app
# COPY . .
# # RUN go mod download
# ENV GO111MODULE=on
# RUN CGO_ENABLED=0 GOOS=linux go build -o ./bankapp 

# FROM alpine:latest

# # RUN go mod download
# # RUN go mod tidy
# COPY --from=build-stage --chown=nonroot:nonroot /go/src/app/assets ./app/assets/
# COPY --from=build-stage --chown=nonroot:nonroot /go/src/app/pkg ./app/pkg/
# COPY --from=build-stage --chown=nonroot:nonroot /go/src/app/views ./app/views/
# COPY --from=build-stage --chown=nonroot:nonroot /go/src/app/bankapp ./app/

# EXPOSE 8080

# ENTRYPOINT ["./app/bankapp", "-p", "8080"]
# ENTRYPOINT ["sleep","1000000000"]