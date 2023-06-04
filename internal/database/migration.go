package database

import (
	"github.com/go-nunu/nunu-layout/internal/model"
	"github.com/go-nunu/nunu-layout/pkg/log"
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
	m.db.AutoMigrate(&model.User{})
	m.log.Logger.Info("AutoMigrate end")
}
