package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"ly-scheduler/src/utils"
	"os"
	"time"
)

const (
	Week       = "week"
	Month      = "month"
	Year       = "year"
	MemberSize = "memberSize"
)

var (
	AllOptions AllConfigOptions
)

type AllConfigOptions struct {
	MainOptions Main `toml:"main"`
}

type Main struct {
	Desc                  string `toml:"desc"`                  // description of the program
	YearMonthDay          string `toml:"yearMonthDay"`          // scheduler start date
	MemberFilePath        string `toml:"memberFilePath"`        // file that contains members
	SchedulePeriod        string `toml:"schedulePeriod"`        // schedule period = week, month, year, memberSize
	ScheduleTableFilePath string `tome:"scheduleTableFilePath"` // file path of saving schedule table
}

func init() {
	// configuration init
	AllOptions = getConfigOptions()
}

// Set AllConfigOptions fields default value.
func (configOptions *AllConfigOptions) setDefaultValues() {
	currentDir, _ := os.Getwd()

	configOptions.MainOptions.Desc = "A scheduler program"
	configOptions.MainOptions.YearMonthDay = utils.GetYYYYMMDDDate()
	configOptions.MainOptions.MemberFilePath = fmt.Sprintf("%s%cmember.xlsx", currentDir, os.PathSeparator)
	configOptions.MainOptions.SchedulePeriod = Week
	configOptions.MainOptions.ScheduleTableFilePath = fmt.Sprintf("%s%c排班表.xlsx", currentDir, os.PathSeparator)
}

func dayTimeByString(yyyyMmDd string) time.Time {
	now, _ := time.Parse("2006-01-02", yyyyMmDd)
	year, month, day := now.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, now.Location())

}

func dayTimeAfterMonth(yyyyMmDd string) time.Time {
	now, _ := time.Parse("2006-01-02", yyyyMmDd)
	year, month, day := now.Date()
	return time.Date(year, month+1, day, 0, 0, 0, 0, now.Location())
}

// Return time slice of every day in month
func dayTimesOfMonth(yyyyMmDd string) []time.Time {
	// dayCounter of month = lastDay - firstDay + 24Hour
	diffHours := dayTimeAfterMonth(yyyyMmDd).Sub(dayTimeByString(yyyyMmDd)).Hours()
	diffHours = diffHours + 24
	dayCounter := int(diffHours / 24)

	var dayTimes []time.Time
	for i := 0; i < dayCounter; i++ {
		dayTimes = append(dayTimes, dayTimeByString(yyyyMmDd).Add(time.Duration(i)*24*time.Hour))
	}
	return dayTimes
}

func dayTimesOfYear(yyyyMmDd string) []time.Time {
	now, _ := time.Parse("2006-01-02", yyyyMmDd)
	year, month, day := now.Date()
	currentDay := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	sameDayAfterYear := time.Date(year+1, month, day, 0, 0, 0, 0, now.Location())
	diffHours := sameDayAfterYear.Sub(currentDay).Hours()

	// days of year = lastDay - firstDay + 24Hour
	diffHours = diffHours + 24
	dayCounter := int(diffHours / 24)

	var dayTimesOfYearSlice []time.Time
	for i := 0; i < dayCounter; i++ {
		dayTimesOfYearSlice = append(dayTimesOfYearSlice, currentDay.Add(time.Duration(i)*24*time.Hour))
	}
	return dayTimesOfYearSlice
}

func dayTimesOfMemberSize(yyyyMmDd string, memberSize int) []time.Time {
	dayTimesOfMonthSlice := dayTimesOfMonth(yyyyMmDd)
	if memberSize <= len(dayTimesOfMonthSlice) {
		return dayTimesOfMonthSlice[0:memberSize]
	}

	dayTimesOfYearSlice := dayTimesOfYear(yyyyMmDd)

	if memberSize > len(dayTimesOfYearSlice) {
		memberSize = len(dayTimesOfYearSlice)
		fmt.Printf("memberSize > dayTimesOfYearSlice, set memberSize = dayTimesOfYearSlice.\n")
	}
	return dayTimesOfYearSlice[0:memberSize]
}

// Get days of schedule period, the result decide by user configuration,
// when schedulePeriod == MemberSize, input parameter will take effect
func (configOptions *AllConfigOptions) SchedulePeriodDays(memberSize int) []time.Time {
	var dayTimes []time.Time

	// select schedule period by config
	switch configOptions.MainOptions.SchedulePeriod {
	case Week:
		dayTimes = dayTimesOfMonth(configOptions.MainOptions.YearMonthDay)[0:7]
	case Month:
		dayTimes = dayTimesOfMonth(configOptions.MainOptions.YearMonthDay)
	case Year:
		dayTimes = dayTimesOfYear(configOptions.MainOptions.YearMonthDay)
	case MemberSize:
		dayTimes = dayTimesOfMemberSize(configOptions.MainOptions.YearMonthDay, memberSize)
	default:
		dayTimes = dayTimesOfMonth(configOptions.MainOptions.YearMonthDay)
	}
	return dayTimes
}

func readConfig(path string) (AllConfigOptions, error) {
	configOptions := AllConfigOptions{}
	configOptions.setDefaultValues()
	if _, err := toml.DecodeFile(path, &configOptions); err != nil {
		return AllConfigOptions{}, err
	}
	return configOptions, nil
}

func getConfigOptions() AllConfigOptions {
	// read config from file
	configFilePaths := []string{"./config.toml", "./src/config.toml", "./conf/config.toml", "../conf/config.toml"}
	found := false
	for _, v := range configFilePaths {
		if config, err := readConfig(v); err == nil {
			AllOptions = config
			found = true
			break
		}
	}
	if !found {
		log.Fatal("cannot find config.toml file. exit program.")
	}
	return AllOptions
}
