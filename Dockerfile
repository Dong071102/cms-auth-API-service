# Bước 1: Sử dụng Go image chính thức làm base image
FROM golang:1.19-alpine

# Bước 2: Cài đặt các thư viện cần thiết (nếu cần)
# Bạn có thể thêm các gói phụ thuộc hệ thống nếu cần thiết cho ứng dụng Go của bạn
RUN apk add --no-cache git

# Bước 3: Đặt thư mục làm việc trong container
WORKDIR /app

# Bước 4: Sao chép tệp go.mod và go.sum (nếu có) vào container
COPY go.mod go.sum ./

# Bước 5: Tải xuống các phụ thuộc của ứng dụng (tạo file go.sum và go.mod)
RUN go mod tidy

# Bước 6: Sao chép toàn bộ mã nguồn của ứng dụng vào container
COPY . .

# Bước 7: Biên dịch ứng dụng Go
RUN go build -o main .

# Bước 8: Expose port 9000
EXPOSE 9000

# Bước 9: Chạy ứng dụng Go trên cổng 9000
CMD ["./main"]
