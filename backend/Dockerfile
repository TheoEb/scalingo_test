FROM golang:1.12-alpine as builder

RUN apk add --no-cache dep git
WORKDIR /go/src/github.com/TheoEb/scalingo_test/backend
COPY Gopkg.lock Gopkg.toml ./
RUN dep ensure -v -vendor-only
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/TheoEb/scalingo_test/backend/server .
CMD ["./server"]
