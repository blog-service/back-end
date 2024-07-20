package jwt

import (
	"back-end/pkg/logger"
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	global     Service
	consoleLog = logger.ConsoleLog()
)

type Service interface {
	GenerateToken(tokenId string, isRefreshToken bool, duration time.Duration) (tokenStr string, err error)
	ValidateToken(token string) (claim *Claims, err error)
}

type service struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func NewJwtService(privateKeyPath, publicKeyPath string) Service {
	if privateKey == nil || publicKey == nil {
		privateKeyByte, err := os.ReadFile(privateKeyPath)
		if err != nil {
			return nil
		}
		publicKeyByte, err := os.ReadFile(publicKeyPath)
		if err != nil {
			return nil
		}
		privateKey, err = parsePrivateKey(privateKeyByte)
		if err != nil {
			return nil
		}
		publicKey, err = parsePublicKey(publicKeyByte)
		if err != nil {
			return nil
		}
	}

	return &service{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

type ClaimsData struct {
	IsAccess bool
}

type Claims struct {
	jwt.RegisteredClaims
	Data ClaimsData
}

func GetGlobal() Service {
	return global
}
