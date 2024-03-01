FROM golang:1.22-alpine3.19 as build
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY ./ ./

RUN GOOS=linux go build -o /todo-app

FROM alpine:3.19.1 AS release
WORKDIR / 
COPY --chmod=0755 --from=build /todo-app ./
COPY --from=build /app/templ/ ./templ/

EXPOSE 8090

ENTRYPOINT [ "/todo-app" ]