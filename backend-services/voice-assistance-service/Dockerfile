FROM golang:1.19 AS build-stage

RUN groupadd appgroup && useradd -g appgroup appuser

WORKDIR /build

COPY go.mod ./

COPY go.sum ./

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o voice-assistance-service .

FROM alpine:3.14 AS final

COPY --from=build-stage /etc/passwd /etc/passwd

COPY --from=build-stage /build/voice-assistance-service /app/voice-assistance-service

USER appuser

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./voice-assistance-service"]
