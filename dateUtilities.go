package Utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const format  = "2006/01/02"
func GetYearMonthDay(date string) (year int,month int,day int,err error) {

	if date=="" {
		date = time.Now().Format(format)
	}

	dateParsedToday:=strings.Split(date,"/")
	year, errParsing := strconv.Atoi(dateParsedToday[0])
	if errParsing != nil {

		return -1,-1,-1, errors.New(ErrorinStringToIntConversion)
	}

	month, errParsing = strconv.Atoi(dateParsedToday[1])
	if errParsing != nil {

		return -1,-1,-1, errors.New(ErrorinStringToIntConversion)
	}
	day, errParsing = strconv.Atoi(dateParsedToday[2])
	if errParsing != nil {

		return -1,-1,-1, errors.New(ErrorinStringToIntConversion)
	}

	return year,month,day,nil
}
