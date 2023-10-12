package server

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
	"time"
)

type Task struct {
	log       *log.Logger
	scheduler *gocron.Scheduler
}

func NewTask(log *log.Logger) *Task {
	return &Task{
		log: log,
	}
}
func (t *Task) Start(ctx context.Context) error {
	// eg: crontab task
	t.scheduler = gocron.NewScheduler(time.UTC)

	_, err := t.scheduler.CronWithSeconds("0/3 * * * * *").Do(func() {
		t.log.Info("I'm a Task1.")
	})
	if err != nil {
		t.log.Error("Task1 error", zap.Error(err))
	}

	_, err = t.scheduler.Every("3s").Do(func() {
		t.log.Info("I'm a Task2.")
	})
	if err != nil {
		t.log.Error("Task2 error", zap.Error(err))
	}

	t.scheduler.StartBlocking()
	return nil
}
func (t *Task) Stop(ctx context.Context) error {
	t.scheduler.Stop()
	t.log.Info("Task stop...")
	return nil
}
