package sid

import (
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/convert"
	"github.com/sony/sonyflake"
)

type Sid struct {
	sf *sonyflake.Sonyflake
}

func NewSid() *Sid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	return &Sid{sf}
}
func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", err
	}
	return convert.IntToBase62(int(id)), nil
}
func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}
