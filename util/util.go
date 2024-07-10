package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get MD5
func Getmd5(msg interface{}) string {
	str, _ := msg.(string)
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

type Myclaims struct {
	UserId primitive.ObjectID `json:"_id"`
	Email  string             `json:"email"`
	jwt.RegisteredClaims
}

var secret = []byte("BAOER-IM-SYS")

// Gettoken
func Gettoken(userid primitive.ObjectID, email string) (string, error) {
	claims := Myclaims{
		UserId: userid,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			Issuer:    "baoer",                                            //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return str, nil
}

// Parsetoken
func Parsetoken(tokenstr string) (*Myclaims, error) {
	myclaim := &Myclaims{}
	claims, err := jwt.ParseWithClaims(tokenstr, myclaim, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analys Token Error:%v", err)
	}
	return myclaim, nil
}

func CreateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(r.Intn(10))
	}
	return code
}
