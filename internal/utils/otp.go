package utils

import (
    "fmt"
    "time"
)

func GenerateOTP() string {
    return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}
