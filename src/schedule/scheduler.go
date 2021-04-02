package schedule

import "ly-scheduler/src/model"

type Scheduler interface {
	DoSchedule(string) []model.MemberRecord
}
