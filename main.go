package main

import (
	"errors"
	"log"

	"github.com/armon/go-socks5"
)

type usernameAuth struct {
	// 保存用户名和密码的映射
	users map[string]string
}

func (a *usernameAuth) Valid(username, password string) bool {
	// 判断用户名和密码是否匹配
	return a.users[username] == password
}

func (a *usernameAuth) Authenticate(username string, password string) error {
	// 使用 Valid 方法判断用户名和密码是否匹配
	if !a.Valid(username, password) {
		return errors.New("invalid username or password")
	}
	return nil
}

func main() {
	// 创建一个保存用户名和密码的映射
	users := map[string]string{
		"garden": "garden",
		"user2":  "password2",
		// ...
	}
	// 创建一个认证器
	authenticator := &usernameAuth{users: users}

	// 创建一个新的 SOCKS5 服务器
	server, err := socks5.New(&socks5.Config{
		// 指定认证器
		Credentials: authenticator,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 监听本地的 1080 端口
	if err := server.ListenAndServe("tcp", "127.0.0.1:1088"); err != nil {
		log.Fatal(err)
	}
}
