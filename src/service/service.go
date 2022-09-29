package service

import (
	"crypto/md5"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/google/wire"
	"maxblog-me-admin/internal/core"
	"time"
)

var ServiceSet = wire.NewSet(
	UserSet,
)

const (
	EmptyStr = ""
)

func genToken(encryptedMobile string, duration time.Duration) (string, string, error) {
	mobile, err := core.RSADecrypt(core.GetPrivateKey(), encryptedMobile)
	if err != nil {
		return EmptyStr, EmptyStr, core.FormatError(202, err)
	}
	j := core.NewJWT()
	token, err := j.GenerateToken(mobile, duration)
	if err != nil {
		return EmptyStr, EmptyStr, core.FormatError(203, err)
	}
	cipherToken, err := core.RSAEncrypt(core.GetPublicKey(), token)
	if err != nil {
		return EmptyStr, EmptyStr, core.FormatError(204, err)
	}
	return cipherToken, mobile, nil
}

func encodePwd(pwd string) (string, string) {
	options := password.Options{
		SaltLen:      16,
		Iterations:   64,
		KeyLen:       16,
		HashFunction: md5.New,
	}
	salt, encodedPwd := password.Encode(pwd, &options)
	return salt, encodedPwd
}
