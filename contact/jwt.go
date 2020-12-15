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
		forgetTime, isForget := IsForgetToken(token)
		// 失效token且失效超过5s
		// 5s是为了
		// 假设A请求发起
		// 发起A后, B请求发起refreshToken
		// 由于不可抗力A在B结束后才请求到服务器, A持有的token此时已forget
		// 给与A 5s缓冲时间, 也就是说refreshToken后5s原token也就是已失效token仍有效
		if isForget && time.Now().Unix()-forgetTime > 5 {
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

	now := time.Now().Unix()
	if RedisClient == nil {
		return
	}
	result := RedisClient.SetNX(
		context.Background(),
		fmt.Sprintf("jwt_forget:%s", token),
		now,
		time.Duration(j.Claim.ExpiresAt-now)*time.Second)
	lib.AssetsError(result.Err())
}

func IsForgetToken(token string) (int64, bool) {
	if RedisClient == nil {
		return 0, false
	}
	cmd := RedisClient.Get(context.Background(), lib.StringJoin("jwt_forget:", token))
	if cmd.Err() != nil {
		return 0, false
	} else {
		if value, err := cmd.Int64(); err != nil {
			return 0, true
		} else {
			return value, true
		}
	}
}
