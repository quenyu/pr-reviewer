FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pr-reviewer ./cmd/server

FROM gcr.io/distroless/base-debian12 AS runner
WORKDIR /app
COPY --from=builder /app/pr-reviewer /app/pr-reviewer
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/pr-reviewer"]
