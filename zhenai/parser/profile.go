package parser

import (
	"regexp"
	"strconv"

	"imooc.com/learngo/crawler/model"

	"imooc.com/learngo/crawler/engine"
)

// age
var ageRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)

// height
var heightRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)

// weight
var weightRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)

// marriage
var marriageRegex = regexp.MustCompile(`<div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

// HoKou
var hoKouRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)

// 职业和教育
var occupyAndEducationRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div></div>`)

// 收入
var incomeRegex = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)

var idUrlRe = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)

// profile
//var profile = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^\\b]+) | ([\d]+)岁 | ([^|]+) | ([^\\b]+) | ([\d]+)cm | ([\d]+-[\d]+)元</div> <div class="actions" data-v-3c42fade>`)

// ParseProfile ParseProfile
func ParseProfile(contents []byte, Url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRegex))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRegex)

	height, err := strconv.Atoi(extractString(contents, heightRegex))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRegex))
	if err == nil {
		profile.Weight = weight
	}
	profile.HoKou = extractString(contents, hoKouRegex)
	matches := occupyAndEducationRegex.FindSubmatch(contents)
	if len(matches) > 2 {
		profile.Occupation = string(matches[1])
		profile.Education = string(matches[2])
	}
	profile.Income = extractString(contents, incomeRegex)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id:      extractString([]byte(Url), idUrlRe),
				Type:    "zhenai",
				Url:     Url,
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, regx *regexp.Regexp) string {
	matches := regx.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}
