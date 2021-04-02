package schedule

import (
	"ly-scheduler/src/config"
	"ly-scheduler/src/model"
	"math/rand"
	"time"
)

var (
	RsScheduler RandomScheduler
)

type RandomScheduler struct {
	MrSlice    []model.MemberRecord
	AllOptions config.AllConfigOptions
}

func shuffle(mrSlice []model.MemberRecord) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(mrSlice), func(i, j int) {
		mrSlice[i], mrSlice[j] = mrSlice[j], mrSlice[i]
	})
}

func init() {
	RsScheduler = RandomScheduler{
		MrSlice:    model.MrSlice,
		AllOptions: config.AllOptions,
	}
}

func (rs *RandomScheduler) DoSchedule(string) []model.ScheduleRecord {
	shuffle(rs.MrSlice)

	validMemberRecordSlice := selectMemberRecordByStatus(rs.MrSlice, 0)
	size := len(validMemberRecordSlice)
	schedulePeriodDayTimes := rs.AllOptions.SchedulePeriodDays(size)

	var srSlice []model.ScheduleRecord // row total size is fixed, equal to member size.
	for index, mr := range validMemberRecordSlice {
		sr := model.ScheduleRecord{}
		sr.TableContent = make(map[string]string)
		sr.Name = mr.MemberName
		for i := 0; i < len(schedulePeriodDayTimes); i++ {
			if i == index || (i-index)%size == 0 {
				sr.TableContent[schedulePeriodDayTimes[i].Format("2006-01-02")] = "Y"
			} else {
				sr.TableContent[schedulePeriodDayTimes[i].Format("2006-01-02")] = ""
			}
		}
		srSlice = append(srSlice, sr)
	}

	invalidMemberRecordSlice := selectMemberRecordWithNotStatus(rs.MrSlice, 0)
	for _, mr := range invalidMemberRecordSlice {
		sr := model.ScheduleRecord{}
		sr.TableContent = make(map[string]string)

		sr.Name = mr.MemberName
		for i := 0; i < len(schedulePeriodDayTimes); i++ {
			sr.TableContent[schedulePeriodDayTimes[i].Format("2006-01-02")] = "-"
		}
		srSlice = append(srSlice, sr)
	}
	return srSlice
}

func selectMemberRecordByStatus(mrSlice []model.MemberRecord, status int64) []model.MemberRecord {
	var wantStatusSlice []model.MemberRecord
	for _, mr := range mrSlice {
		if mr.SchedulingStatus == status {
			wantStatusSlice = append(wantStatusSlice, mr)
		}
	}
	return wantStatusSlice
}

func selectMemberRecordWithNotStatus(mrSlice []model.MemberRecord, status int64) []model.MemberRecord {
	var wantStatusSlice []model.MemberRecord
	for _, mr := range mrSlice {
		if mr.SchedulingStatus != status {
			wantStatusSlice = append(wantStatusSlice, mr)
		}
	}
	return wantStatusSlice
}
