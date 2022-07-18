package utils

import(
	"encoding/json"
	"io/ioutil"
	// "net/http"
	"strings"
	"os"
	"log"
	"fmt"
	"net/http"
	// "regexp"
	// "gopkg.in/Knetic/govaluate.v2"
	"github.com/tmdvs/Go-Emoji-Utils"
	"ehe.com/publicsonar/defs"

)

func TextCleaner(str string) string{
	// re, err := regexp.Compile(`[^]`)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// str1 = re.ReplaceAllString(str1, " ")

	str = emoji.RemoveAll(str)
	str = strings.ToLower(str)
	str = strings.TrimSuffix(str, "\n")
	return str
}

func ReadCases() []defs.Cases{
	var cases []defs.Cases

    // Open our jsonFile
    jsonFile, err := os.Open("storage/cases.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    // read our opened xmlFile as a byte array.
    file, _ := ioutil.ReadAll(jsonFile)

    err = json.Unmarshal(file, &cases)
    if err != nil {
        log.Fatal(err)
    }
    return cases
}

func ReadMessages() []string{
	var arr []string

	// Open our jsonFile
	jsonFile, err := os.Open("storage/messages.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	file, _ := ioutil.ReadAll(jsonFile)

	// err = json.Unmarshal(file, &cases)
	err = json.Unmarshal(file, &arr)
	if err != nil {
		log.Fatal(err)
	}
	return arr
}


func ParseBody(r *http.Request, x interface {}){
if body, err := ioutil.ReadAll(r.Body); err == nil {
	if err := json.Unmarshal([]byte(body), x); err != nil {
		return 
	}
}
}

