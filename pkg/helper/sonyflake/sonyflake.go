package sonyflake

import (
	"github.com/sony/sonyflake"
	"time"
)

func NewSonyflake() *sonyflake.Sonyflake {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC),
		MachineID:      nil,
		CheckMachineID: nil,
	})
	if sf == nil {
		panic("sonyflake not created")
	}
	return sf
}
