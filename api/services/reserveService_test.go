package services

import (
	"fmt"
	"testing"

	"github.com/go-easylog/el"
)

func TestGetReserve(t *testing.T) {
	reserve := getReserve()
	fmt.Println(reserve)
}

func TestPersistReserve(t *testing.T) {
	el.SetLogLevel(el.TRACE)
	reserve := persistReserve()
	fmt.Println(reserve)
}
