package core

import (
	"crypto/md5"
	"crypto/rsa"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var ctx *Context
var once sync.Once

func init() {
	once.Do(func() {
		ctx = &Context{}
	})
}

func GetInstanceOfContext() *Context {
	return ctx
}

type Context struct {
	Upstream        Upstream
	Downstream      Downstream
	JWTSecret       string
	PwdEncodingOpts *password.Options
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	PublicKeyStr    string
}

type Upstream struct {
	MaxblogFEAdmin AddressHttp
}

type Downstream struct {
	MaxblogBEUser Address
	MaxblogBEDemo Address
}

type Address struct {
	Host string
	Port int
}

type AddressHttp struct {
	Protocol string
	Domain   string
	Host     string
	Port     int
	Secure   bool
}

func GetUpstreamAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Upstream.MaxblogFEAdmin.Host, ctx.Upstream.MaxblogFEAdmin.Port)
}

func GetUpstreamDomain() string {
	return fmt.Sprintf("%s://%s", ctx.Upstream.MaxblogFEAdmin.Protocol, ctx.Upstream.MaxblogFEAdmin.Domain)
}

func GetUpstreamSecure() bool {
	return ctx.Upstream.MaxblogFEAdmin.Secure
}

func GetDownstreamMaxblogBEUserAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Downstream.MaxblogBEUser.Host, ctx.Downstream.MaxblogBEUser.Port)
}

func GetDownstreamMaxblogBEDemoAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Downstream.MaxblogBEDemo.Host, ctx.Downstream.MaxblogBEDemo.Port)
}

func GetProjectPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	indexWithoutFileName := strings.LastIndex(path, string(os.PathSeparator))
	indexWithoutLastPath := strings.LastIndex(path[:indexWithoutFileName], string(os.PathSeparator))
	return strings.Replace(path[:indexWithoutLastPath], "\\", "/", -1)
}

func GetPublicKey() *rsa.PublicKey {
	return ctx.PublicKey
}

func GetPublicKeyStr() string {
	return ctx.PublicKeyStr
}

func GetPrivateKey() *rsa.PrivateKey {
	return ctx.PrivateKey
}

func SetKeys() {
	prk, puk, _ := GenRsaKeyPair(2048)
	ctx.PublicKey = puk
	ctx.PrivateKey = prk
	publicKeyStr, _ := PublicKeyToString()
	ctx.PublicKeyStr = publicKeyStr
}

func SetJWTSecret(jwtSecret string) {
	ctx.JWTSecret = jwtSecret
}

func SetPwdEncodingOpts() {
	ctx.PwdEncodingOpts = &password.Options{
		SaltLen:      16,
		Iterations:   64,
		KeyLen:       16,
		HashFunction: md5.New,
	}
}

func GetEncodedPwd(pwd string) (string, string) {
	salt, encodedPwd := password.Encode(pwd, ctx.PwdEncodingOpts)
	return salt, encodedPwd
}

func VerifyEncodedPwd(pwdHeldRaw string, salt string, pwdTarget string) bool {
	return password.Verify(pwdHeldRaw, salt, pwdTarget, ctx.PwdEncodingOpts)
}
