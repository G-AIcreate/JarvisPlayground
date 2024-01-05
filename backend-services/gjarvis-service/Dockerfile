# Build Stage
FROM golang:1.19 AS build-stage

# Create appuser.
RUN groupadd appgroup && useradd -g appgroup appuser
# RUN adduser -D -g '' appuser
# -D: 这个选项告诉 adduser 命令在创建用户时不分配密码，因为你的应用可能不需要登录或使用 shell。
# -g '': 这个选项设置新创建用户的初始登录组或改变现有用户的登录组。在这里，它用一个空字符串 '' 来表示没有初始登录组。

WORKDIR /build

# Use https://goproxy.cn as the Go proxy.
# ENV GOPROXY=https://goproxy.cn

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
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o gjarvis-service .

# Final Stage
FROM alpine:3.14 AS final
# FROM gcr.io/distroless/base-debian11 AS build-release-stage

# Import the user and group files from the builder.
COPY --from=build-stage /etc/passwd /etc/passwd

# Copy the binary from builder stage to final stage.
COPY --from=build-stage /build/gjarvis-service /app/gjarvis-service

# Use an unprivileged user.
USER appuser

WORKDIR /app

# Expose port 8080 to the outside world.
EXPOSE 8081

# Command to run the executable.
ENTRYPOINT ["./gjarvis-service"]
