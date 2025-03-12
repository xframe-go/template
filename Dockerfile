# 第一阶段：构建阶段（Builder）
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 拷贝代码
COPY . .

# 构建可执行文件，注意静态编译参数
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o xonic .

# 第二阶段：生成最终极简镜像
FROM scratch

# 拷贝构建好的二进制文件
COPY --from=builder /app/xonic /xonic

# 设置入口命令
ENTRYPOINT ["/xonic", "serve"]