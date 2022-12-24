FROM golang:1.19.1 as builder

WORKDIR /app

COPY . ./

RUN go build -o /api ./cmd/api/

FROM golang:1.19.1 as dev

WORKDIR /app

COPY . ./

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "build/api/.air.toml"]

FROM gcr.io/distroless/static-debian11 as prod

WORKDIR /

COPY --from=builder /api /api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/api" ]
