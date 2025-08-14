FROM golang:1.21.5-buster as builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

RUN go run github.com/steebchen/prisma-client-go prefetch

COPY ./ ./

RUN go run github.com/steebchen/prisma-client-go generate

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app ./cmd/server

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /workspace/app /app

ENTRYPOINT ["/app"]