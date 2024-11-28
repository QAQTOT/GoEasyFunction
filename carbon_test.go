package GoEasyFunction

import (
	"go_easy_funcition/easy_request"
	"testing"
)

func TestCarbon(t *testing.T) {

	type Person struct {
		Name    string
		Age     int
		Address string
	}
	form, err := easy_request.PostJson("https://www.baidu.com", "/", Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	})

	t.Log(form, err)
	//t.Log(quick_func.SubString("111212ads", 0, -1))
	//
	//t.Log(carbon.Now().GetDateTimeString())
	//t.Log(carbon.Now().GetDateString())
	//t.Log(carbon.Now().GetTimeString())
	//
	//t.Log(carbon.Now().GetUnixTimeStamp())
	//t.Log(carbon.Now().GetUnixMicroTimeStamp())
	//t.Log(carbon.Now().GetUnixNanoTimeStamp())

}
