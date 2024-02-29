FROM golang:1.22-alpine3.19 as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

RUN GOOS=linux go build -o /todo-app

FROM alpine:3.19.1 AS release
COPY --chmod=0755 --from=build /todo-app .

EXPOSE 8090

ENTRYPOINT [ "/todo-app" ]