package cron

import (
	"time"

	"github.com/robfig/cron/v3"
)

type Adapter interface {
	GetNextRun(expr string) (*time.Time, error)
}

type adapterImpl struct{}

func NewAdapter() Adapter {
	return &adapterImpl{}
}

func (r *adapterImpl) GetNextRun(expr string) (*time.Time, error) {
	sched, err := cron.ParseStandard(expr)
	if err != nil {
		return nil, err
	}

	next := sched.Next(time.Now())

	return &next, nil
}
