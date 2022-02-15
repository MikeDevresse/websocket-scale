FROM golang:1.17-alpine as build

RUN mkdir /app

COPY ./ /app/

WORKDIR /app

RUN go mod download; \
    go build -o main ./

CMD ["/app/main"]

FROM busybox AS prod

COPY --from=build /app/main ./

CMD ["/main"]