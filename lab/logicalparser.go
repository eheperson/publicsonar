package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
	"os"
	"strings"
	"math/rand"
)

type Cases struct{
    CaseId    int `json:"case_id"`
    Query string `json:"query"`
    Queries string `json:"queries"`
}

func readCases() []Cases{
	var cases []Cases

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

const (
	OperatorAND = "AND"
	OperatorOR = "OR"
	LeftBrackets = "("
	RightBrackets = ")"
)


type Group struct{
	Begin int
	End int
}

type Term struct{

}

func groupParser(query string) []Group{
	var bracketsLocations []int
	var bracketsTypes []string

	for pos, char := range query {
		// fmt.Printf(" %c - %d\n", char, pos)
		if string(char) == LeftBrackets {
			bracketsLocations = append(bracketsLocations, pos)
			bracketsTypes = append(bracketsTypes, LeftBrackets)
		} else if string(char) == RightBrackets{
			bracketsLocations = append(bracketsLocations, pos)
			bracketsTypes = append(bracketsTypes, RightBrackets)
		}
	}
	fmt.Println(query)
	// fmt.Println(bracketsLocations)
	// fmt.Println(bracketsTypes)

	var groups []Group
	var tempGroup Group
	var leftBracketsLocations []int


	for pos, char := range bracketsTypes{
		if char == LeftBrackets{
			leftBracketsLocations = append(leftBracketsLocations, bracketsLocations[pos])
		}
		if char == RightBrackets{
			tempGroup = Group{
				Begin:leftBracketsLocations[len(leftBracketsLocations)-1], 
				End:bracketsLocations[pos],
			}
			groups = append(groups, tempGroup)
			leftBracketsLocations = leftBracketsLocations[:len(leftBracketsLocations)-1]
		}
	}
	return groups
}


// RandomString generates a random string of length n
func RandomString(n int) string {
	alphabet := "abcdefghijklmnoprstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}	

	return sb.String()
}

func FilterOR(message string, groups []Group, query string){
	var flagOR bool
	var group string

	for _, value := range groups{
		flagOR = false
		group = query[value.Begin:value.End+1]
	
		// fmt.Println(group)
		// fmt.Println(value.Begin, value.End+1)

		ORCounter := strings.Count(group, OperatorOR)

		if ORCounter == 1 {
			splitedGroup := strings.Split(group, OperatorOR)

			splitedGroup0 := strings.Trim(splitedGroup[0], "()")
			splitedGroup0 = strings.TrimSpace(splitedGroup0)
			splitedGroup0 = strings.Trim(splitedGroup[0], "()")
			splitedGroup0 = strings.TrimSpace(splitedGroup0)
			splitedGroup1 := strings.Trim(splitedGroup[1], "()")
			splitedGroup1 = strings.TrimSpace(splitedGroup1)
			splitedGroup1 = strings.Trim(splitedGroup[1], "()")
			splitedGroup1 = strings.TrimSpace(splitedGroup1)

			fmt.Println("splitedGroup0 : ", splitedGroup0)
			fmt.Println("splitedGroup0 : ", len(splitedGroup0))
			fmt.Println("splitedGroup1 : ", splitedGroup1)
			fmt.Println("splitedGroup1 : ", len(splitedGroup1))
			
			result1 := strings.Contains(message, splitedGroup0)
			fmt.Println("result1: ", result1)
			if result1 {
				flagOR = true
			}

			result2 := strings.Contains(message, splitedGroup1)
			fmt.Println("result2: ", result2)
			if result2{
				flagOR = true
			}
			if flagOR {
				tempStr := RandomString(len(group))
				fmt.Println("group: ", group)
				query = strings.Replace(query, group, tempStr, 1)
				message = message + tempStr
			}
		}
	}
	fmt.Println("new: ", query)
}


// func main() {
// 	cases := readCases()
// 	messagee := "Real Madrid vs Eintracht Frankfurt • Super Cup • FIFA 22 PS5 Realistic Sliders  • 4K 60fps Gameplay - #fifa22 #ps5 #fifa23   \n\nReal Madrid vs Eintracht Frankfurt • Super Cup • FIFA 22 PS5 Realistic Sliders  • 4K 60fps Gameplay\n\n#istephan\n#eintrachtfrankfurt  \n#realmadrid \n#mbappe\n#messi\n#neymar\n#benzema"
	
// 	groupss := groupParser(cases[3].Queries)
// 	// fmt.Println("Eeeeeee: ",strings.Contains(messagee, "messi"))

// 	queryy := cases[3].Queries
// 	FilterOR(messagee, groupss, queryy)

// 	fmt.Println(groupParser(cases[0].Query))
// 	fmt.Println(groupParser(cases[1].Query))
// 	fmt.Println(groupParser(cases[2].Query))
// 	fmt.Println(groupParser(cases[3].Queries))
// 	fmt.Println(RandomString())
// }