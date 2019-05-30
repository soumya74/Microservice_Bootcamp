# First Stage======================
FROM golang:1.8

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/gorilla/mux
RUN make build_linux

# Second Stage=====================
FROM alpine
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
EXPOSE 10001

COPY --from=0 /app /app
CMD ["./app"]
