package utils

import (
	"regexp"
	"strings"
)

func GetDomain(url string) string {
	reg := regexp.MustCompile("https?://")
	url = string(reg.ReplaceAll([]byte(url), []byte("")))
	return strings.Split(url, "/")[0]
}

func ParseMajestic(str string) []map[string]string {
	arr := strings.Split(str, "\n")

	var keyMap = map[int]string{}
	var parseArr = []map[string]string{}
	var reg = regexp.MustCompile(`\s{2,}`)

	for i, v := range arr {
		l := reg.ReplaceAllString(v, "\t")
		line := strings.Split(l, "\t")
		parse := map[string]string{}
		for k, v2 := range line {
			if i == 0 {
				keyMap[k] = v2
			} else {
				parse[keyMap[k]] = v2
			}
		}
		if i != 0 {
			parseArr = append(parseArr, parse)
		}
	}
	return parseArr
}
