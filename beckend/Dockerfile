FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
# Подтягиваем нужный toolchain (если go.mod требует более свежий)
ENV GOTOOLCHAIN=auto
RUN go mod download

COPY beckend ./beckend

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./beckend/cmd/api

FROM gcr.io/distroless/base-debian12
COPY --from=builder /bin/app /bin/app

ENV APP_HTTP_ADDR=:8080
EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/bin/app"]
