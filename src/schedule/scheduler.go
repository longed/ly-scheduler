package schedule

type Schedule interface {
	DoSchedule(string) string
}
