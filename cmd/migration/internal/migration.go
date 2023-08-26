package internal

import (
	"github.com/go-nunu/nunu-layout-advanced/internal/model"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Migrate struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMigrate(db *gorm.DB, log *log.Logger) *Migrate {
	return &Migrate{
		db:  db,
		log: log,
	}
}
func (m *Migrate) Run() {
	if err := m.db.AutoMigrate(&model.User{}); err != nil {
		m.log.Error("user migrate error", zap.Error(err))
		return
	}
	m.log.Info("AutoMigrate end")
}
