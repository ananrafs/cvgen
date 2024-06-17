package cvgen

import (
	"fmt"
	"time"
)

type Period struct {
	StartTime time.Time
	EndTime   time.Time
}

func (p Period) String() string {
	endPeriod := p.EndTime.Format("Jan 2006")
	if p.EndTime.After(time.Now()) {
		endPeriod = "Present"
	}
	return fmt.Sprintf("%s – %s", p.StartTime.Format("Jan 2006"), endPeriod)
}
