package services

import (
	"apiServer/api/entity"
	"testing"
	"time"

	"github.com/go-easylog/el"
)

func TestPersistRoom(t *testing.T) {
	el.SetLogLevel(el.TRACE)
	room := entity.Room{}
	room.RoomName = "aaaa"
	room.Address = "千葉県船橋市"
	room.Type = entity.COMMUTING
	room.Status = entity.USING
	room.Note = "aaaaaa"
	room.Created = time.Now()
	room.Modified = time.Now()
	persistEntity("rooms", room)
}
