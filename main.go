package main

import (
	"fmt"
	go_module "github.com/yohanapriyandi/first-go-module"
)

// pastikan kita sudah melakukan perintah go get nama-module yang ingin di install
func main() {
	fmt.Println(go_module.SayHello())
}
