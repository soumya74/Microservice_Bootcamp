# First Stage======================
FROM golang:1.8

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/gorilla/mux
RUN go get github.com/mediocregopher/radix.v2
RUN go get -u github.com/go-redis/redis
RUN make build_linux

# Second Stage=====================
FROM alpine
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
EXPOSE 8080

COPY --from=0 /app /app
CMD ["./app"]
