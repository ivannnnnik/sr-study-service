FROM golang:1.25-alpine AS builder
RUN apk add --no-cache gcc musl-dev pkgconf
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o server ./cmd/server 


FROM alpine:3.19
COPY --from=builder /app/server /server
CMD ["/server"]