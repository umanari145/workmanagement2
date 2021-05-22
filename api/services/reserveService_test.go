package services

import (
	"apiServer/api/entity"
	"fmt"
	"testing"
	"time"

	"github.com/go-easylog/el"
)

func TestGetReserve(t *testing.T) {
	reserve := getReserve()
	fmt.Println(reserve)
}

func TestPersistReserve(t *testing.T) {
	el.SetLogLevel(el.TRACE)
	reserve := entity.Reserve{}
	reserve.RoomID = 1
	reserve.UserID = 1
	reserve.StartReserveDate = time.Now()
	reserve.EndReserveDate = time.Now()
	reserve.Created = time.Now()
	reserve.Modified = time.Now()
	persistReserve(reserve)
}
