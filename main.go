package main
import(
    "fmt"
	"errors"
	"time"
	"io/ioutil"
	"encoding/json"
	"jaxf-github.fanatics.corp/BirthdayTest/feature"
	"jaxf-github.fanatics.corp/BirthdayTest/util"
	)
func main() {
	fmt.Println("Parsing File on path ./InputDate/persondetails.txt for Input Data")
	inputData ,err:=parseInputFile()
	if err != nil{
         panic(any(util.ErrorUnmarshalInputDataFile))
	}

    fmt.Println("Input Data is ...")
	fmt.Println(inputData)

	pDataStore:=feature.PeopleDataStore{}
    	errInit:=pDataStore.Init(&inputData)
    	if errInit != nil{
        		panic(any(errInit.Error()))
        	}
        var oper feature.Operations = pDataStore
    	ans,err:= oper.FindTodayPeopleBirthday(time.Now())

    	if err != nil{
    		fmt.Println(err.Error())
    	}else {
    		fmt.Println("Persons whose DOB  today are")
    		fmt.Println(*ans)
    	}

}

func parseInputFile() (personDetails feature.ListInputPeopleDetails, err error ) {

dataFile, err := ioutil.ReadFile(util.InputFilePath)
if err != nil{
return nil,errors.New(util.InputFileReadError)
}
data := []byte(dataFile)
pDetails := feature.ListInputPeopleDetails{}
err = json.Unmarshal(data, &pDetails)
if err !=nil{
return nil,err
}
return  pDetails,nil
}
