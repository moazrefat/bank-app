FROM golang:1.21.5 AS build-stage

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY * ./

RUN CGO_ENABLED=0 GOOS=darwin go build -o /bankapp


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /bankapp /bankapp

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/bankapp", "-p", "8080"]