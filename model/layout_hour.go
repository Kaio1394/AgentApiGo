package model

import (
	"fmt"
	"time"
)

var Layout_date string = "2006-01-02 15:04:05"
var Sysdate time.Time = GetSysdate()

func ConvertDate(dateStr string, layout string) (time.Time, error) {
	var date time.Time
	convert, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Erro ao converter a string para time.Time:", err)
		return date, err
	}
	return convert, nil
}

func GetSysdate() time.Time {
	return time.Now()
}

func AddDays(date *time.Time, qtyDays int) time.Time {
	return date.Add(time.Duration(qtyDays) * 24 * time.Hour)
}
