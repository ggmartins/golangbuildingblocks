FROM golang:latest as builder
WORKDIR /app
COPY src/minibank/go.mod src/minibank/go.sum ./
RUN go mod download
COPY src/minibank/main.go .
RUN go build -o main main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["/main"]
