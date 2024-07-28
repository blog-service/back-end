package jwt

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func parsePrivateKey(keyByte []byte) (*rsa.PrivateKey, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyByte)
	if err != nil {
		consoleLog.Error().Err(err).Str("func", "ParsePrivateKey-jwt.ParseRSAPrivateKeyFromPEM").Msg("jwtPkg")
		return nil, err
	}
	return key, nil
}

func parsePublicKey(keyByte []byte) (*rsa.PublicKey, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyByte)
	if err != nil {
		consoleLog.Error().Err(err).Str("func", "ParsePrivateKey-jwt.ParseRSAPrivateKeyFromPEM").Msg("jwtPkg")
		return nil, err
	}
	return key, nil
}

func (s *service) GenerateToken(tokenId string, isRefreshToken bool, duration time.Duration) (tokenStr string, err error) {
	payload := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			ID:        tokenId,
		},
		Data: ClaimsData{
			IsAccess: isRefreshToken,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	return token.SignedString(s.privateKey)
}

func (s *service) ValidateToken(token string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.New("token validate failed")
		}
		return s.publicKey, nil
	}

	payload := new(Claims)
	if _, err := jwt.ParseWithClaims(token, payload, keyFunc); err != nil {
		consoleLog.Error().Err(err).Str("func", "ValidateToken-jwt.ParseWithClaims").Msg("jwtPkg")
		return nil, err
	}
	return payload, nil
}
