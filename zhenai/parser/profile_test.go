package parser

import (
	"io/ioutil"
	"testing"

	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_data_test.html")
	if err != nil {
		panic(err)
	}
	url := "http://m.zhenai.com/u/1992679628"
	result := ParseProfile(contents, url, "小耳朵")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profile := result.Items[0]
	expected := engine.Item{
		Url:  url,
		Id:   "",
		Type: "",
		Payload: model.Profile{
			Name: "小耳朵",
			// Gender:     "女",
			Age:        31,
			Height:     175,
			Weight:     63,
			Income:     "3-5千",
			Marriage:   "离异",
			Education:  "大专",
			Occupation: "人事专员",
			HoKou:      "阿坝阿坝县",
			// XingZuo:    "天蝎座",
			// House:   "和家人同住",
			// Car:     "未买车",
		},
	}
	if profile != expected {
		t.Errorf("user profile is get failed , expected user profile is : %v, but got is %v", expected, profile)
	}
}
