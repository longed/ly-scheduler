package schedule

import (
	"fmt"
	"ly-scheduler/src/config"
	"ly-scheduler/src/model"
	"testing"
)

func TestRandomScheduler_DoSchedule(t *testing.T) {
	main := config.Main{
		Desc:           "",
		YearMonthDay:   "2021-04-01",
		MemberFilePath: "",
		SchedulePeriod: "",
	}
	all := config.AllConfigOptions{MainOptions: main}
	var mrSlice []model.MemberRecord
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "ABC",
		SchedulingStatus:       "是",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "EFG",
		SchedulingStatus:       "是",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "HIJ",
		SchedulingStatus:       "否",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	rs := RandomScheduler{
		MrSlice:    mrSlice,
		AllOptions: all,
	}

	srSlice := rs.DoSchedule("")
	fmt.Println(len(srSlice))
}

func TestSelector(t *testing.T) {
	var mrSlice []model.MemberRecord
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "ABC",
		SchedulingStatus:       "是",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "EFG",
		SchedulingStatus:       "是",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	mrSlice = append(mrSlice, model.MemberRecord{
		MemberName:             "HIJ",
		SchedulingStatus:       "否",
		Reason:                 "",
		SpecificWorkingWeekday: "",
	})
	selectMemberRecord(mrSlice, "是", statusEqualSelector)
}
