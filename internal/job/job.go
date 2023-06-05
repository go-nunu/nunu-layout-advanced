package job

import (
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type Job struct {
	db  *gorm.DB
	log *log.Logger
}

func NewJob(db *gorm.DB, log *log.Logger) *Job {
	return &Job{
		db:  db,
		log: log,
	}
}
func (j *Job) Run() {
	c := cron.New()
	var err error
	err = c.AddFunc("0/3 * * * * *", func() {
		j.log.Info("I'm a Task.")
	})

	if err != nil {
		fmt.Println(err)
	}

	c.Start()
	select {}
}
