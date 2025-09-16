# مرحله اول: ساخت باینری
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o otp-login cmd/main.go

# مرحله دوم: اجرای باینری
FROM debian:bullseye-slim

WORKDIR /app

# ✅ کپی باینری
COPY --from=builder /app/otp-login .

# ✅ کپی مستندات Swagger
COPY --from=builder /app/docs ./docs

# ✅ تنظیم متغیر محیطی برای JWT
ENV JWT_SECRET=mysecretkey

EXPOSE 8080

CMD ["./otp-login"]
