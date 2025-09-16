# 📲 OTP Login Service

سیستم احراز هویت ساده با OTP و JWT، نوشته‌شده با Go و Gin.

---

## 📦 ویژگی‌ها

- دریافت OTP با شماره تلفن  
- اعتبارسنجی OTP و صدور JWT  
- ذخیره‌سازی OTP به‌صورت موقت در حافظه (in-memory)  
- ثبت‌نام خودکار کاربر جدید در صورت عدم وجود  
- محدودسازی درخواست OTP (rate limiting)  
- مدیریت کاربران با API  
- مستندسازی کامل با Swagger UI  
- اجرای سریع با Docker  

---

## 🧑‍💻 اجرای محلی

```bash
export JWT_SECRET=your-secret-key
go run cmd/main.go
```

> توجه: برای اجرای محلی باید متغیر محیطی `JWT_SECRET` ست شود.  
> Swagger UI پس از اجرا در آدرس زیر در دسترس خواهد بود:  
> `http://localhost:8080/swagger/index.html`

## 🚀 اجرای پروژه در داکر

```bash
docker build -t otp-login-service .
docker run -p 8080:8080 otp-login-service
```

> Swagger UI در آدرس زیر در دسترس خواهد بود:  
> `http://localhost:8080/swagger/index.html`
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

- OTP فقط به‌صورت موقت در حافظه نگه‌داری می‌شود و پس از ۲ دقیقه منقضی می‌شود  
- کلید JWT به‌صورت متغیر محیطی (`JWT_SECRET`) تنظیم می‌شود  
- این پروژه برای تست و آموزش طراحی شده و آماده‌ی توسعه برای محیط واقعی است  
---

## 🛠 توسعه‌های پیشنهادی

\- اتصال به MongoDB یا Redis برای ذخیره‌سازی پایدار  
\- اضافه کردن middleware برای اعتبارسنجی JWT  
\- ساخت لایه‌ی کاربران و نقش‌ها  
\- تست‌های واحد با `go test`  
\- پیاده‌سازی ارسال واقعی OTP از طریق SMS یا پیام‌رسان‌ها  