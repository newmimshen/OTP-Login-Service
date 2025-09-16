package storage

import (
    "sync"
    "time"
)

// ساختار کاربر
type User struct {
    Phone        string    `json:"phone"`
    RegisteredAt time.Time `json:"registered_at"`
}

// ساختار ذخیره‌سازی کاربران
type UserStorage struct {
    data map[string]User
    mu   sync.RWMutex
}

// سازنده‌ی UserStorage
func NewUserStorage() *UserStorage {
    return &UserStorage{
        data: make(map[string]User),
    }
}

// بررسی وجود کاربر
func (s *UserStorage) Exists(phone string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    _, exists := s.data[phone]
    return exists
}

// افزودن کاربر جدید
func (s *UserStorage) Add(phone string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[phone] = User{
        Phone:        phone,
        RegisteredAt: time.Now(),
    }
}

// دریافت اطلاعات یک کاربر خاص
func (s *UserStorage) Get(phone string) (User, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    user, exists := s.data[phone]
    return user, exists
}

// لیست همه‌ی کاربران
func (s *UserStorage) List() []User {
    s.mu.RLock()
    defer s.mu.RUnlock()
    users := make([]User, 0, len(s.data))
    for _, u := range s.data {
        users = append(users, u)
    }
    return users
}
