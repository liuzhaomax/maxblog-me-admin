package core

import (
	"crypto/rsa"
	"fmt"
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
	Upstream     Upstream
	Downstream   Downstream
	JWTSecret    string
	PrivateKey   *rsa.PrivateKey
	PublicKey    *rsa.PublicKey
	PublicKeyStr string
}

type Upstream struct {
	MaxblogFEAdmin Address
}

type Downstream struct {
	MaxblogBEUser Address
	MaxblogBEDemo Address
}

type Address struct {
	Host string
	Port int
}

func SetUpstreamAddr(host string, port int) {
	ctx.Upstream.MaxblogFEAdmin.Host = host
	ctx.Upstream.MaxblogFEAdmin.Port = port
}

func GetUpstreamAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Upstream.MaxblogFEAdmin.Host, ctx.Upstream.MaxblogFEAdmin.Port)
}

func SetDownstreamBEUserAddr(host string, port int) {
	ctx.Downstream.MaxblogBEUser.Host = host
	ctx.Downstream.MaxblogBEUser.Port = port
}

func GetDownstreamBEUserAddr() string {
	return fmt.Sprintf("%s:%d", ctx.Downstream.MaxblogBEUser.Host, ctx.Downstream.MaxblogBEUser.Port)
}

func SetDownstreamBEDemoAddr(host string, port int) {
	ctx.Downstream.MaxblogBEDemo.Host = host
	ctx.Downstream.MaxblogBEDemo.Port = port
}

func GetDownstreamBEDemoAddr() string {
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
	return GetInstanceOfContext().PublicKey
}

func GetPublicKeyStr() string {
	return GetInstanceOfContext().PublicKeyStr
}

func GetPrivateKey() *rsa.PrivateKey {
	return GetInstanceOfContext().PrivateKey
}
