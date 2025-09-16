package storage

import (
    "sync"
    "time"
)

// RateLimiter محدودکننده‌ی تعداد درخواست OTP برای هر شماره تلفن
type RateLimiter struct {
    requests map[string][]time.Time
    mu       sync.Mutex
}

// NewRateLimiter سازنده‌ی RateLimiter
func NewRateLimiter() *RateLimiter {
    return &RateLimiter{
        requests: make(map[string][]time.Time),
    }
}

// Allow بررسی می‌کند که آیا شماره تلفن مجاز به دریافت OTP جدید هست یا نه
// محدودیت: حداکثر ۳ درخواست در ۱۰ دقیقه
func (rl *RateLimiter) Allow(phone string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()
    windowStart := now.Add(-10 * time.Minute)

    // گرفتن لیست درخواست‌های قبلی
    requests := rl.requests[phone]

    // فیلتر کردن درخواست‌های معتبر در پنجره‌ی ۱۰ دقیقه‌ای
    filtered := make([]time.Time, 0, len(requests))
    for _, t := range requests {
        if t.After(windowStart) {
            filtered = append(filtered, t)
        }
    }

    // بررسی محدودیت
    if len(filtered) >= 3 {
        return false
    }

    // ثبت درخواست جدید
    filtered = append(filtered, now)
    rl.requests[phone] = filtered

    return true
}
