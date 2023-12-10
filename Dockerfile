# 使用官方 Go 镜像作为构建环境
FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod ./

# 复制源代码
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用 Alpine Linux
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从 builder 镜像中复制可执行文件
COPY --from=builder /app/main .

# 运行应用程序
CMD ["./main"]