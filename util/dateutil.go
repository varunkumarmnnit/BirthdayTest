package util

import (
	"time"
	"errors"
)

func IsLeapYear(date time.Time)  bool {
	var isLeapYear bool = false
     if !date.IsZero()  && date.Year() %4 ==0 {
		 isLeapYear= true
	 }
	 return isLeapYear
}

func CheckValidDate(date time.Time) error{
 if int(date.Year()) < 0 || date.IsZero(){
 return errors.New(InValidDate)
 }
 return nil
 }

