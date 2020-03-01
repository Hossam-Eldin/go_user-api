package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow : return time
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString : get time for functions as string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}
