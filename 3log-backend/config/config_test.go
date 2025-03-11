package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// 需要修改 AddConfigPath 路径为 ‘.’
	Init()
	fmt.Printf("%+v", C)
}
