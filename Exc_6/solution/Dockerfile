FROM golang:1.25 AS builder
WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y bash git && \
    go install github.com/swaggo/swag/cmd/swag@latest

RUN bash /app/scripts/build-application.sh

RUN bash /app/scripts/build-openapi-docs.sh

FROM alpine AS run
WORKDIR /app

COPY --from=builder /app/ordersystem /app/ordersystem
COPY --from=builder /app/docs /app/docs

EXPOSE 3000

CMD ["/app/ordersystem"]
