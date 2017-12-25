package slackbot

import (
	"io/ioutil"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func sanitizeKey(key string) string {
	return strings.TrimSpace(key)
}

// GETAPIToken reads the key from the file
func GetAPIToken(keyFile string) string {

	data, err := ioutil.ReadFile(keyFile)
	checkErr(err)

	//sanitize the string
	key := sanitizeKey(string(data))
	return key
}
