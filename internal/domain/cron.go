package domain

import "time"

type CronExpression struct {
	Description   string
	NextExecution time.Time
}

func (c *CronExpression) FormattedNextExecution() string {
	return c.NextExecution.Format("2006-01-02 15:04:05")
}
