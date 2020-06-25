package contact

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	lib "github.com/maxiloEmmmm/go-tool"
	"time"
)

type Claim struct {
	jwt.StandardClaims
}

type JwtLib struct {
	Claim
	Secret []byte
}

func JwtNew() *JwtLib {
	now := time.Now()

	instance := &JwtLib{}

	instance.Secret = []byte(Config.Jwt.Secret)

	instance.Claim.StandardClaims = jwt.StandardClaims{
		Audience:  "front",
		Issuer:    "XJ-LA",
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
	}
	return instance
}

func (j *JwtLib) SetPrimaryKey(value string) *JwtLib {
	j.Claim.Id = value
	return j
}

func (j *JwtLib) SetExpiresAt(value int64) *JwtLib {
	j.Claim.ExpiresAt = value
	return j
}

func (j *JwtLib) SetExpiresAtWeek() *JwtLib {
	j.Claim.ExpiresAt = time.Now().Add(7 * 24 * time.Hour).Unix()
	return j
}

func (j *JwtLib) SetExpiresAtDay(value int64) *JwtLib {
	j.Claim.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	return j
}

func (j *JwtLib) SetExpiresAt2Hour(value int64) *JwtLib {
	j.Claim.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return j
}

func (j *JwtLib) SetSubInfo(value string) *JwtLib {
	j.Claim.Subject = value
	return j
}

func (j *JwtLib) SetSecret(secret string) *JwtLib {
	j.Secret = []byte(secret)
	return j
}

func (j *JwtLib) GenerateToken() (token string) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claim).SignedString(j.Secret)
	lib.AssetsError(err)
	return token
}

func (j *JwtLib) ParseToken(token string) (err error) {
	tokenObj, err := jwt.ParseWithClaims(token, &j.Claim, func(token *jwt.Token) (i interface{}, err error) {
		return j.Secret, nil
	})

	if err != nil {
		return err
	}

	if _, ok := tokenObj.Claims.(*Claim); ok && tokenObj.Valid {
		if IsForgetToken(token) {
			return jwt.NewValidationError("失效token", 0)
		}

		return nil
	} else if !ok {
		return errors.New("jwt类型不匹配")
	} else {
		return err
	}
}

func (j *JwtLib) RefreshToken(addTime time.Duration) (token string) {
	j.ForgetToken()
	j.Claim.ExpiresAt = time.Now().Add(addTime).Unix()
	token = j.GenerateToken()
	return token
}

func (j *JwtLib) ForgetToken() {
	token := j.GenerateToken()

	result := RedisClient.SetNX(
		context.Background(),
		fmt.Sprintf("jwt_forget:%s", token),
		1,
		time.Duration(j.Claim.ExpiresAt-time.Now().Unix())*time.Second)
	lib.AssetsError(result.Err())
}

func IsForgetToken(token string) bool {
	return RedisClient.Exists(context.Background(), fmt.Sprintf("jwt_forget:%s", token)).Val() == 1
}
