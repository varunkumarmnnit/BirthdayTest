package feature

import (
	"strings"
	"testing"
	"time"
	"errors"
   "jaxf-github.fanatics.corp/BirthdayTest/util"

)

func TestFindTodayPeopleBirthdayNoData(t *testing.T) {
	pDetails:= ListInputPeopleDetails{}
	pDataStore:= PeopleDataStore{}
    gotErr:= pDataStore.Init(&pDetails)
    want:= util.ErrorNoInputDataProvided

   if gotErr.Error() != want{
	t.Errorf(" Wanted %s as error but got %s",want,gotErr.Error())
  }
}

func TestFindTodayPeopleBirthdayIncorrectDataFormat(t *testing.T) {

	 errorTestCases :=[]struct{
		description string
		pDataStore PeopleDataStore
		input ListInputPeopleDetails
		want string
	}{
		 {
			 description: "DOB Missing",pDataStore: PeopleDataStore{},input: [][]string{{"Varun"}},want: "Incorrect Data format",

		 },
		 {
			 description: "Only One Name Given",pDataStore: PeopleDataStore{},input: [][]string{{"Varun","2022/05/10"}},want: "Incorrect Data format for person",

		 },
		 {
			 description: "Incorrect DOB Format",pDataStore: PeopleDataStore{},input: [][]string{{"Varun","Kumar","05/10/2022"}},want: "Incorrect DOB Format",

		 },


	 }
	for _,tt:=range errorTestCases{
		t.Run(tt.description, func(t *testing.T) {
			got:=tt.pDataStore.Init(&tt.input)
			if !strings.Contains(got.Error(),tt.want){
				t.Errorf(" Wanted %s as error but got %s",tt.want,got.Error())

			}
		})

	}

}

func TestFindTodayPeopleBirthday(t *testing.T) {

	    errorTestCases :=[]struct{
		description string
		inputDate time.Time
		pDetails ListInputPeopleDetails
		lenExpOutPutList int
		errExp error
		nameExpec string
	}{
		 {
			 description: "No output found",inputDate:time.Now().AddDate(0,1,0),pDetails:[][]string{{"Varun","Kumar",time.Now().Format(util.DateFormat)}}, lenExpOutPutList:0,errExp:errors.New(util.NoResultFound),nameExpec:"" ,

		 },
		 {
			 description: "Two Output found",inputDate:time.Now(),pDetails:[][]string{{"David","Norbert",time.Now().Format(util.DateFormat)},{"Varun","Kumar",time.Now().Format(util.DateFormat)} },lenExpOutPutList:2,errExp:nil,nameExpec:"Norbert" ,

		 },
		 {
         	description: "Invalid Date Entered",inputDate:time.Date(-1962,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),pDetails:[][]string{{"Varun","Kumar","1964/02/29"}},lenExpOutPutList:0,errExp:errors.New(util.InValidDate),nameExpec:"",

         },
		 {
			 description: "Non Leap Year Condition",inputDate:time.Date(1962,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),pDetails:[][]string{{"Varun","Kumar","1964/02/29"},{"David","Steves",time.Now().Format(util.DateFormat)}},lenExpOutPutList:1,errExp:nil,nameExpec:"Kumar",

		 },
		 {
         	description: "Non Leap Year Condition With Leap Year Input",inputDate:time.Date(1962,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),pDetails:[][]string{{"Varun","Kumar","1964/02/29"},
                                                                                                                                                                                                                                       			{"David","Steves","1964/02/28"},
                                                                                                                                                                                                                                       			                                                                                                                                                                                                                               			{"David","Charles","1963/02/28"},},lenExpOutPutList:3,errExp:nil,nameExpec:"Kumar",

         },
		 {
         	description: " Leap Year Condition",inputDate:time.Date(1964,time.February,28,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),pDetails:[][]string{{"Varun","Kumar","1964/02/29"},
                                                                                                                                                                                                               			{"David","Steves","1964/02/28"},
                                                                                                                                                                                                               			{"David","Charles","1963/02/28"},},lenExpOutPutList:2,errExp:nil,nameExpec:"Steves" ,
         },
         {
                  	description: " Leap Year Condition Vary Input ",inputDate:time.Date(1964,time.February,29,time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Nanosecond(),time.Now().Location()),pDetails:[][]string{{"Varun","Kumar","1964/02/29"},
                                                                                                                                                                                                                        			{"David","Steves","1964/02/28"},
                                                                                                                                                                                                                        			{"David","Charles","1963/02/28"},},lenExpOutPutList:1,errExp:nil,nameExpec:"Kumar" ,
                  },

    }

    	for _,tt:=range errorTestCases{
    		t.Run(tt.description, func(t *testing.T) {
            pDetails:=ListInputPeopleDetails{}
            pDataStore:=PeopleDataStore{}
            pDetails = tt.pDetails
            pDataStore.Init(&pDetails)
            ans,qErr:=pDataStore.FindTodayPeopleBirthday(tt.inputDate)
            if qErr != nil && qErr.Error() != tt.errExp.Error(){
            t.Errorf(" Wanted %s as error but got %s",tt.errExp.Error(),qErr.Error())
            }else if qErr != nil  && tt.errExp == nil{
            t.Errorf("Not expecting any error for the given input")
            }else{
            if ans != nil && (len(*ans)!=tt.lenExpOutPutList || (*ans)[0][1]!= tt.nameExpec){
            t.Errorf("Expected output is not matching with thr output returned by FindTodayPeopleBirthday method ")
            		}

    		}})

    	}
    }
