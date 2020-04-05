from golang:alpine as builder


run apk update && apk add --no-cache git

workdir /app

copy go.mod go.sum ./

run go mod download

copy . . 

run CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world .

FROM alpine:latest
run apk --no-cache add ca-certificates 

workdir /root/

copy --from=builder /app/hello-world .
copy --from=builder /app/templates /root/templates

EXPOSE 8080

cmd [ "./hello-world"]