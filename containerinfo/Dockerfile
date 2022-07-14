FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /out/containerinfo

EXPOSE 8000

FROM alpine:3.16.0

COPY --from=builder /out/containerinfo /containerinfo

CMD ["/containerinfo"]
