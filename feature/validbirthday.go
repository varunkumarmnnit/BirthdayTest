package feature
import (
	"time"
	"errors"
	"fmt"
	"jaxf-github.fanatics.corp/BirthdayTest/util"
)

// This will contain the input data parsed from file or command
type ListInputPeopleDetails [][]string

// This will contain the output result
type ListTodayPeopleBirthday [][]string

//Interface that will have the method to show People having birthday today based on input data
type Operations interface {
	FindTodayPeopleBirthday(todayDate time.Time) (*ListTodayPeopleBirthday,error)
}

//This is people detail in specific format .Better for readability and doing calculation based on date
type People struct {
	firstName,lastName string
	birthday  time.Time
}

//This is the store  that will contain the all  person detail in the specific format
type PeopleDataStore struct {
	allPeople []People
}

//This is the constructor which will parse the people detail from generic 2D list to specific structure format.
func (p *PeopleDataStore) Init(pDetails *ListInputPeopleDetails)error   {

//Validating  that their should be some input present to calculate Birthday Operation
if len(*pDetails) == 0 {
		return  errors.New(util.ErrorNoInputDataProvided)
	}

	for _, eachPersonDetail := range *pDetails {
	   //Checking that user must pass three values as input ie first name ,last name and DOB
		if len(eachPersonDetail) != 3 {

			return  errors.New(fmt.Sprintf("Incorrect Data format for person %s.It Should be in  the format Last Name,First Name ,DOB in format YYYY/MM/DD",eachPersonDetail))
		}
		birthday, err := time.Parse(util.DateFormat, eachPersonDetail[2])

		//The DOB should be correct and with the specific format
		if err != nil {
			return  errors.New(fmt.Sprintf("Incorrect DOB Format for person %s.It should be in format YYYY/MM/DD",eachPersonDetail))
		}
		p.allPeople=append(p.allPeople,People{firstName: eachPersonDetail[0], lastName: eachPersonDetail[1], birthday: birthday})

	}
	return nil
}

//This function calculates people whose DOB is today .We passes todayDate as parameter to increase the unit testing coverage and checking all the conditions.
func (p PeopleDataStore) FindTodayPeopleBirthday(todayDate time.Time) ( *ListTodayPeopleBirthday, error) {
var ans ListTodayPeopleBirthday
     dateErr:=util.CheckValidDate(todayDate)
     if dateErr != nil {
     return nil,dateErr

     }

  isTodayLeapYear := util.IsLeapYear(todayDate)
	for _, eachPersonDetail := range p.allPeople {
	   //Checking if todays month is Friday
		if todayDate.Month() == time.February {
			if isTodayLeapYear {
				validateBirthday(todayDate, eachPersonDetail, &ans)
			} else {
			   //This is the case where person born on Leap year on 29th Feb but todays date is not leap years and date is 28.
				if eachPersonDetail.birthday.Day() == util.FebNoDaysWithoutLeapYear+1 && todayDate.Day() == util.FebNoDaysWithoutLeapYear {
					ans = append(ans, []string{eachPersonDetail.firstName, eachPersonDetail.lastName, eachPersonDetail.birthday.String()})
				} else {
					validateBirthday(todayDate, eachPersonDetail, &ans)
				}
			}
		} else {
			validateBirthday(todayDate, eachPersonDetail, &ans)
		}

	}
    if len(ans)==0{
		return nil,errors.New(util.NoResultFound)
	}
	return &ans, nil
}

func validateBirthday(todaysDate time.Time,personDetail People, ans *ListTodayPeopleBirthday) {
	if todaysDate.Day() == personDetail.birthday.Day() && todaysDate.Month() == personDetail.birthday.Month() {
		*ans = append(*ans, []string{personDetail.firstName, personDetail.lastName, personDetail.birthday.Format(util.DateFormat)})

	}
}
