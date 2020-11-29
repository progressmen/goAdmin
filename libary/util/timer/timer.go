package timer

import (
	"strconv"
	"time"
)

var (
	DefaultDateFormat       = "2006-01-02"
	DefaultDatetimeFormat   = "2006-01-02 15:04:05"
	FormatYYMMDDHH          = "2006010215"
	DefaultDateFormatZh     = "2006年1月2日"
	DefaultDatetimeFormatZh = "2006年1月2日 15时04分05秒"
)

// 秒
func SeTime() int64 {
	return time.Now().Unix()
}

// 毫秒
func MsTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// 微秒
func UsTime() int64 {
	return time.Now().UnixNano() / 1e3
}

// 纳秒
func NsTime() int64 {
	return time.Now().UnixNano()
}

func MsFrontTime(ms int64) string {
	ms = ms / 1e3
	return time.Unix(ms, 0).Format(DefaultDatetimeFormat)
}

func DateFormat(format string) string {
	now := time.Now()
	return now.Format(format)
}

func TimestampMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func DatetimeString() string {
	return time.Now().Format(DefaultDatetimeFormat)
}
func DateString() string {
	return time.Now().Format(DefaultDateFormat)
}

//func DateFormat(t time.Time) string {
//	return t.Format(DefaultDateFormat)
//}

func DatetimeFormat(t time.Time) string {
	return t.Format(DefaultDatetimeFormat)
}

func DatetimeParse(t string) (time.Time, error) {
	return time.Parse(DefaultDatetimeFormat, t)
}

func DatetimeParseUnix(t string) int64 {
	var timeInfo, _ = time.Parse(DefaultDatetimeFormat, t)
	return timeInfo.Unix()
}

func DatetimeParseLocal(t string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(DefaultDatetimeFormat, t, loc)
}

//func ToDate(t time.Time) time.Time {
//	t, _ = time.Parse(DefaultDateFormat, DateFormat(t))
//	return t
//}

/**
获取2006-01-02 15:04:05.000 ms时间
*/
func DatetimeMS() string {
	t := DatetimeString()
	ms := strconv.FormatInt(TimestampMS(), 10)[10:]
	return t + "." + ms
}

/**
获取当天0点毫秒单位的时间戳
*/
func TodayDateTimeMS() int64 {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	return startTime.UnixNano() / 1e6
}

/**
获取某天0点毫秒单位的时间戳
*/
func GetDayDateTimeMS(year, month, day int) int64 {
	currentTime := time.Now()
	startTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, currentTime.Location())
	return startTime.UnixNano() / 1e6
}

/**
 * @func 把时间化为更易读的日期和时间
 * @param times  时间
 * @param timeLayout  时间模板
 * @return times  根据时间模板返回时间格式
 */
func ConversionTime(times string, timeLayout string) string {
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", times, loc)
	timestamp := tmp.Unix()
	times = time.Unix(timestamp, 0).Format(timeLayout)
	return times
}

// 当前周几
func WeekDay() int {
	t := time.Now()
	return int(t.Weekday())
}

func ConversionTimeMSToDate(timeMs int64, timeLayout string) string {
	return time.Unix(timeMs/1000, 0).Format(timeLayout)

}
