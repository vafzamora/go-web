FROM golang:1.22-alpine3.19 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

RUN apk add --update gcc musl-dev \
    && CGO_ENABLED=1 GOOS=linux go build -o /go-web \
    && rm *.db -rf

FROM alpine:3.19.1 AS release
COPY --chmod=0755 --from=build /go-web .

EXPOSE 8080

ENTRYPOINT [ "/go-web" ]