package cron

import (
	"github.com/mathcale/what-the-cron/internal/domain"
	"github.com/mathcale/what-the-cron/internal/pkg/cron"
)

type CronUseCase interface {
	Execute(expression string) (domain.CronExpression, error)
}

type ucImpl struct {
	cronAdapter cron.Adapter
	humanizer   *CronHumanizer
}

func NewCronUseCase(adapter cron.Adapter) CronUseCase {
	return &ucImpl{
		cronAdapter: adapter,
		humanizer:   NewCronHumanizer(),
	}
}

func (uc *ucImpl) Execute(expression string) (domain.CronExpression, error) {
	human, err := uc.humanizer.Humanize(expression)
	if err != nil {
		return domain.CronExpression{}, err
	}

	next, err := uc.cronAdapter.GetNextRun(expression)
	if err != nil {
		return domain.CronExpression{}, err
	}

	return domain.CronExpression{
		Description:   human,
		NextExecution: *next,
	}, nil
}
