package sid

import (
	"github.com/go-nunu/nunu-layout-advanced/pkg/helper/convert"
	"github.com/pkg/errors"
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
	// 生成分布式ID
	id, err := s.sf.NextID()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate sonyflake ID")
	}
	// 将ID转换为字符串
	return convert.IntToBase62(int(id)), nil
}
func (s Sid) GenUint64() (uint64, error) {
	// 生成分布式ID
	return s.sf.NextID()
}
