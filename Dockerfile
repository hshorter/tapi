FROM damianoneill/golang-alpine-builder AS builder
WORKDIR $GOPATH/src/github.com/damianoneill/tapi
COPY . ./
RUN go build -o /app github.com/damianoneill/tapi/cmd/tapi-server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 8080
COPY --from=builder /app ./
CMD ["./app", "--host=0.0.0.0","--port=8080"]