package helpers

import (
	"server/app/pkg/constants"
	"time"
)

func defaultTimeLocation() *time.Location {
	return time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)
}

func ParseStringToStartsOfDay(date string) (*string, error) {
	if parseDate, err := time.Parse(constants.YYYYMMDD_DateFormat, date); err != nil {
		return nil, err
	} else {
		convertedDate := parseDate.Format(constants.YYYYMMDD_HHMMSS_DateTimeFormat)
		return &convertedDate, nil
	}
}
