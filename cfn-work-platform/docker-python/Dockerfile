# 第一阶段：构建 Go 二进制
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

# 启用 Go Modules 并构建
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /cfn-work-platform

# 第二阶段：构建最终镜像（包含 Python）
FROM python:3.11-slim

# 从 builder 阶段复制 Go 二进制文件
COPY --from=builder /cfn-work-platform /usr/local/bin/cfn-work-platform

# 设置容器启动命令
CMD ["/usr/local/bin/cfn-work-platform"]