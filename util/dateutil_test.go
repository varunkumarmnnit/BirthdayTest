package util

import (
	"testing"
	"time"
	"errors"

)

func TestIsLeapYear(t *testing.T) {

	 testCases :=[]struct{
    		description string
    		input time.Time
    		want bool
    	}{
    		 {
    			 description: "Validating 1964 for Leap Year",input:time.Date(1964,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),want: true,

    		 },
    		 {
    			 description: "Validating 1951 for Leap Year",input:time.Date(1951,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),want: false,

    		 },

    	 }
    	for _,tt:=range testCases{
    		t.Run(tt.description, func(t *testing.T) {
    			got:=IsLeapYear(tt.input)
    			if got != tt.want{
    				t.Errorf(" Wanted %v as error but got %v",tt.want,got)

    			}
    		})

    	}
}


func TestValidDate(t *testing.T) {

	 testCases :=[]struct{
    		description string
    		input time.Time
    		want error
    	}{
    		 {
    			 description: "InValid Year",input:time.Date(-200,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),want: errors.New(InValidDate),

    		 },
    		 {
    			 description: "Valid Date",input:time.Date(1951,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),want: nil,

    		 },

    	 }
    	for _,tt:=range testCases{
    		t.Run(tt.description, func(t *testing.T) {
    			got:=CheckValidDate(tt.input)
    			if got == nil && tt.want != nil{
    			t.Errorf(" Wanted %v as error but got %v",tt.want.Error(),got)
    			}
    			if got != nil && got.Error() != tt.want.Error(){
    				t.Errorf(" Wanted %v as error but got %v",tt.want.Error(),got.Error())

    			}
    		})

    	}
}