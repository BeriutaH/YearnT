package factory

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// HashPassword 哈希密码，有点慢
func HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

// CheckPassword 检查密码
func CheckPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// DjangoEncrypt 加密
func DjangoEncrypt(password string, sl string) string {
	pwd := []byte(password)
	salt := []byte(sl)
	iterations := 120000
	digest := sha256.New
	dk := pbkdf2.Key(pwd, salt, iterations, 32, digest)
	str := base64.StdEncoding.EncodeToString(dk)
	return "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
}

// GetRandom 获取盐
func GetRandom() []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	deStr := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 12; i++ {
		result = append(result, deStr[r.Intn(len(deStr))])
	}
	return result
}

// DjangoCheckPassword 检查密码
func DjangoCheckPassword(account string, password string) bool {
	sl := strings.Split(account, "$")[2]
	checkPasswordToken := DjangoEncrypt(password, sl)
	if account == checkPasswordToken {
		return true
	} else {
		return false
	}
}
