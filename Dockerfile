FROM golang:1.21.5 AS build-stage

WORKDIR /app
COPY . .

RUN GOOS=linux go build -o /bankapp main.go

# FROM gcr.io/distroless/static-debian11 AS build-release-stage
FROM alpine:latest AS build-release-stage
WORKDIR /

COPY --from=build-stage /bankapp /bankapp

EXPOSE 8080

ENTRYPOINT ["sleep","10000000000"]

# FROM golang:1.21.5

# WORKDIR /
# COPY . .

# RUN GOOS=linux go build -gcflags "all=-N -l" -o /bankapp

# EXPOSE 8080

# ENTRYPOINT ["/bankapp", "-p", "8080"]