FROM golang:1.13-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

COPY . .

RUN go build \
  -a \
  -ldflags "-s -w -extldflags 'static'" \
  -installsuffix cgo \
  -tags netgo \
  -mod vendor \
  -o /bin/malice \
  ./server/main.go




FROM gcr.io/distroless/static
COPY --from=builder /bin/malice /bin/malice
ENTRYPOINT ["/bin/malice"]
