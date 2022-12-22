FROM golang:1.19.1 as builder

WORKDIR /app

COPY . ./

RUN go build -o /sse ./cmd/sse/

FROM golang:1.19.1 as dev

WORKDIR /app
COPY . ./

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "build/sse/.air.toml"]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /sse /sse

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/sse" ]
