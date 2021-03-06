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

	validMemberRecordSlice := selectMemberRecord(rs.MrSlice, model.ChineseYes, statusEqualSelector)
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

	invalidMemberRecordSlice := selectMemberRecord(rs.MrSlice, model.ChineseYes, statusNotEqualSelector)
	for _, mr := range invalidMemberRecordSlice {
		sr := model.ScheduleRecord{}
		sr.TableContent = make(map[string]string)

		sr.Name = mr.MemberName
		for i := 0; i < len(schedulePeriodDayTimes); i++ {
			sr.TableContent[schedulePeriodDayTimes[i].Format("2006-01-02")] = mr.Reason
		}
		srSlice = append(srSlice, sr)
	}
	return srSlice
}

type statusSelectFunc func(string, string) bool

func selectMemberRecord(mrSlice []model.MemberRecord, status string, statusSelector statusSelectFunc) []model.MemberRecord {
	var wantStatusSlice []model.MemberRecord
	for _, mr := range mrSlice {
		if statusSelector(mr.SchedulingStatus, status) {
			wantStatusSlice = append(wantStatusSlice, mr)
		}
	}
	return wantStatusSlice
}

func statusEqualSelector(leftStatus, rightStatus string) bool {
	return leftStatus == rightStatus
}

func statusNotEqualSelector(leftStatus, rightStatus string) bool {
	return leftStatus != rightStatus
}
