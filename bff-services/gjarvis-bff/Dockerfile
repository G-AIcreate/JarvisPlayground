# Build Stage
FROM golang:1.19 AS build-stage

# Create appuser.
RUN groupadd appgroup && useradd -g appgroup appuser

WORKDIR /build

# 安装 protobuf 工具
RUN apt-get update && apt-get install -y protobuf-compiler && \
    go get google.golang.org/protobuf/cmd/protoc-gen-go && \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# 复制 .proto 文件并编译生成 Go 源代码
COPY ./infrastructure/grpc/gjarvis-proto/gjarvis_service.proto ./
RUN protoc -I ./ gjarvis_service.proto --go_out=plugins=grpc:./

# Copy go mod and sum files.
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. They will be cached if the go.mod and go.sum files are not changed.
RUN go mod download -x

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app for a Linux environment and disable CGO.
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o gjarvis-bff .

# Final Stage
FROM alpine:3.14 AS final

# Import the user and group files from the builder.
COPY --from=build-stage /etc/passwd /etc/passwd

# Copy the binary from builder stage to final stage.
COPY --from=build-stage /build/gjarvis-bff /app/gjarvis-bff

# Use an unprivileged user.
USER appuser

WORKDIR /app

# Expose port 8081 (or whatever port you use for your BFF) to the outside world.
EXPOSE 8080

# Command to run the executable.
ENTRYPOINT ["./gjarvis-bff"]
