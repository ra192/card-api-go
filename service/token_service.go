package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
)

var hmacSecret = []byte("UCnmDHn9QS+GqLR5Gkyfw00fykPgW8R9b9uALi4xHEA=")

func CreateToken(merchantId uint, secret string) (string, error) {
	Merchant, err := GetActiveMerchantById(merchantId)
	if err != nil {
		return "", err
	}

	if Merchant.Secret != sha256Hash(secret) {
		return "", errors.New("secret is not valid")
	}

	claims := &jwt.StandardClaims{Subject: strconv.FormatUint(uint64(merchantId),10)}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(hmacSecret)
}

func ValidateToken(token string) (uint, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if err != nil {
		return 0, err
	}
	if !tokenObj.Valid {
		return 0, errors.New("token is not valid")
	}

	claims := tokenObj.Claims.(*jwt.StandardClaims)

	id, err := strconv.ParseUint(claims.Subject, 10, 64)

	return uint(id), err
}

func sha256Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
