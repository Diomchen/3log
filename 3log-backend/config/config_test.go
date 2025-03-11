package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	Init()
	fmt.Printf("%+v", C)
}
