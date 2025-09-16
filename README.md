# 📲 OTP Login Service

سیستم احراز هویت ساده با OTP و JWT، نوشته‌شده با Go و Gin.

---

## 📦 ویژگی‌ها

- دریافت OTP با شماره تلفن  
- اعتبارسنجی OTP و صدور JWT  
- ذخیره‌سازی OTP به‌صورت محلی (in-memory)  
- ساختار ماژولار با هندلرها و لایه‌ی ذخیره‌سازی  
- تست‌پذیر با REST Client  
- اجرای پروژه با Docker  

---

## 🚀 اجرای پروژه

```bash
docker build -t otp-login-service .
docker run -p 8080:8080 otp-login-service
```

---

## 🧪 تست‌ها

فایل‌های تست موجود:

- `test-request.http` → دریافت OTP  
- `test-verify.http` → اعتبارسنجی موفق  
- `test-verify-invalid.http` → اعتبارسنجی ناموفق  

> قابل اجرا با REST Client یا ابزارهای مشابه مثل Thunder Client یا Postman

---

## 📁 ساختار پروژه

```plaintext
internal/
├── handlers/
│   ├── request_otp.go
│   ├── verify_otp.go
│   ├── users.go
├── storage/
│   └── otp.go
cmd/
└── main.go
docs/
└── swagger.json
tests/
├── test-rate-limit.http
├── test-request.http
├── test-verify.http
└── test-verify-invalid.http
```

---

## 📮 مسیرهای API

### `POST /auth/request-otp`

```json
{
  "phone": "09123456789"
}
```

پاسخ:

```json
{
  "message": "OTP ارسال شد (در کنسول چاپ شده)"
}
```

---

### `POST /auth/verify-otp`

```json
{
  "phone": "09123456789",
  "code": "کدی که در کنسول چاپ شده"
}
```

پاسخ موفق:

```json
{
  "message": "ورود موفق",
  "token": "..."
}
```

---

### `GET /users`

پارامترهای اختیاری:

- `page` → شماره صفحه  
- `limit` → تعداد در هر صفحه  
- `search` → جستجو بر اساس شماره تلفن  

---

### `GET /users/{phone}`

دریافت اطلاعات یک کاربر خاص بر اساس شماره تلفن

---

## 🔐 نکات امنیتی

- OTP فقط به‌صورت موقت در حافظه نگه‌داری می‌شود  
- کلید JWT به‌صورت ثابت در کد تعریف شده (`your-secret-key`)  
  → پیشنهاد می‌شود از فایل `.env` خوانده شود  
- این پروژه برای تست و آموزش طراحی شده و آماده‌ی توسعه برای محیط واقعی است

---

## 🛠 توسعه‌های پیشنهادی

- اتصال به MongoDB یا Redis برای ذخیره‌سازی پایدار  
- اضافه کردن middleware برای اعتبارسنجی JWT  
- ساخت لایه‌ی کاربران و نقش‌ها  
- محدودسازی ارسال OTP (rate limiting)  
- اضافه کردن Swagger UI برای مستندات API  
- تست‌های واحد با `go test`
