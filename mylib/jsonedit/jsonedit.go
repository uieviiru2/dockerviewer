package jsonedit

import (
	"regexp"
	"strings"
)

/////////////////////////////////////////////////////////////////////////
//json編集用
/////////////////////////////////////////////////////////////////////////
//jsonデータ同士を接続
func Con(base, add string) string {
	return base + "," + add
}

//最後に{}でくくるために使用
func End(base string) string {
	return "{ " + base + " }"
}

// "aaa": "bbb" といったシンプルなデータを作成
func Val(key, value string) string {
	return "\"" + key + "\": \"" + value + "\""
}

//Keys
func Key(key, value string) string {
	return "\"" + key + "\": {" + value + "}"
}

func On(key, value string) string {
	return "\"" + key + "\": " + value
}

func List(key, value string) string {
	return "\"" + key + "\": [" + value + "]"
}

/*
func JsonObj(base, objName string) string {
	return "\"" + objName + "\": { " + base + " }"
}
*/

func Split(docment, compile, arrName string) string {
	output := ""
	//for i, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(docment), -1) {
	for i, v := range regexp.MustCompile(compile).Split(string(docment), -1) {
		if v == "" {
			continue
		}
		if i == 0 {
			output = v
		} else {
			output = output + "," + v
		}
	}
	return "\"" + arrName + "\": [" + output + "]"
}

//両サイドのダブルクォート削除用
func Split2(docment, compile, arrName string) string {
	output := ""
	//for i, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(docment), -1) {
	for i, v := range regexp.MustCompile(compile).Split(string(docment), -1) {
		if len(v) == 0 {
			continue
		}
		if i == 0 {
			output = v[1 : len(v)-1]
		} else {
			output = output + "," + StripQ(v)
		}
	}
	return "\"" + arrName + "\": [" + output + "]"
}

func StripQ(str string) string {
	return str[1 : len(str)-1]
}

func JsonEscape(val string) string {
	return strings.Replace(val, "\n", "\\n", -1)
	//resEsc, _ := json.Marshal([]byte(val))
	//fmt.Println(string(resEsc))
	//return string(resEsc)
}
