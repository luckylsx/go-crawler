package persist

import (
	"context"
	"encoding/json"
	"testing"

	"imooc.com/learngo/crawler/engine"

	"github.com/olivere/elastic"

	"imooc.com/learngo/crawler/model"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "",
		Type: "",
		Id:   "",
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// save expected item
	err = save("dating_profile", client, expected)
	if err != nil {
		panic(err)
	}

	// fetch saved item
	resp, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	// t.Logf("%s", *resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	t.Logf("%+v", actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	// verify result
	if actual != expected {
		t.Errorf("got %v;expected %v", actual, expected)
	}
}
