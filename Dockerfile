# 使用的基礎映像檔
FROM golang:1.22.1

# 設定工作目錄
WORKDIR /app

# 將您的 Go 應用程式的依賴項目複製到 Docker 映像檔中
COPY go.mod .
COPY go.sum .

# 下載所有依賴項目
RUN go mod download

# 將您的 Go 應用程式的源碼複製到 Docker 映像檔中
COPY . .

# 建立您的 Go 應用程式
RUN go build -o main .

# 定義容器執行的指令
CMD ["./main"]