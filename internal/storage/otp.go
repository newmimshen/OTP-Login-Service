package storage

import "time"

type OTPEntry struct {
    Code      string
    ExpiresAt time.Time
}

var otpStore = make(map[string]OTPEntry)

func SaveOTP(phone, code string, expiresAt time.Time) {
    otpStore[phone] = OTPEntry{Code: code, ExpiresAt: expiresAt}
}

func GetOTP(phone string) (OTPEntry, bool) {
    entry, exists := otpStore[phone]
    return entry, exists
}

func DeleteOTP(phone string) {
    delete(otpStore, phone)
}
