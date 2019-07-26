FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shiiip-vessel

COPY . .

RUN GOPROXY=https://goproxy.io go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shiiip-vessel


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shiiip-vessel/shiiip-vessel .

CMD ["./shiiip-vessel"]