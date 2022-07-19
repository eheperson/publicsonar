package classifier

import (
	"strings"
	"log"
	"encoding/json"
    "io/ioutil"
	"os"
	"gopkg.in/Knetic/govaluate.v2"
	"ehe.com/publicsonar/utils"
	"ehe.com/publicsonar/defs"
	// "fmt"
)

func messageTokenizer(message string) []string{
	return strings.Fields(message)
}

func singleQueryCheck(message string,  query string) bool {
	var ifQueryExistsFlag bool = false
	tokenizedMessage := messageTokenizer(message)

	for i := range tokenizedMessage{
		if strings.Compare(query, tokenizedMessage[i]) == 0 {
			ifQueryExistsFlag = true
		}
		// fmt.Println(tokenizedMessage[i], " ---- ", ifQueryExistsFlag)
	}

	return ifQueryExistsFlag
}

func QueryTokenizer(query string) []string{
	ORCounter := strings.Count(query, "OR")
	ANDCounter := strings.Count(query, "AND")
	LBracketCounter := strings.Count(query, "(")
	RBracketCounter := strings.Count(query, ")")
	
	query = strings.Replace(query, "OR", "XX", ORCounter)
	query = strings.Replace(query, "AND", "XX", ANDCounter)
	query = strings.Replace(query, "(", "", LBracketCounter)
	query = strings.Replace(query, ")", "", RBracketCounter)

	// tokenizedQuery := strings.Fields(query)
	tokenizedQuery := strings.Split(query, "XX")
	for i := range tokenizedQuery {
		tokenizedQuery[i] = strings.TrimSpace(tokenizedQuery[i])
	}
	// fmt.Println(strings.Join(tokenizedQuery[:], ","))
	return tokenizedQuery
}

func QueryRefactor(query, message string) string{
	// originalQuery := query

	var queryTokenLength int
	message = utils.TextCleaner(message)
	tokenizedQuery := QueryTokenizer(query)

	for _, val := range tokenizedQuery{
		queryTokenLength = len(strings.Fields(val))
		// fmt.Println("len : ", queryTokenLength, "token : ", val)

		if queryTokenLength == 1 {
			singleQueryFlag := singleQueryCheck(message, val)

			if singleQueryFlag{
				query = strings.Replace(query, val, "true" , 1)
			} else {
				query = strings.Replace(query, val, "false" , 1)
			}

		} else {
			if strings.Contains(message, val){
				query = strings.Replace(query, val, "true" , 1)
			} else {
				query = strings.Replace(query, val, "false" , 1)
			}
		}

	}
	ORCounter := strings.Count(query, "OR")
	query = strings.Replace(query, "OR", "||", ORCounter)
	
	ANDCounter := strings.Count(query, "AND")
	query = strings.Replace(query, "AND", "&&", ANDCounter)

	SpaceCounter := strings.Count(query, " ")
	query = strings.Replace(query, " ", "", SpaceCounter)
	query = "(" + query + ")"
	return query
}

func MessageClassifier(message string, cases []defs.Cases) []int{
	var query string
	var caseArr []int

	for j := range cases {
		// fmt.Println(cases[j])
		if len(cases[j].Query) != 0{
			query = cases[j].Query
		}
		if len(cases[j].Queries) != 0{
			query = cases[j].Queries
		}

		refactoredQuery := QueryRefactor(query, message)
		// expression := binaryTree(refactoredQuery)
	
		evaluableExp, err := govaluate.NewEvaluableExpression(refactoredQuery)
		if err != nil {
			log.Fatal(err)
		}
		result, err := evaluableExp.Evaluate(nil)
		if err != nil {
			log.Fatal(err)
		}
		if result == true{
		// fmt.Printf("%v -- %T \n",result, result)
			caseArr = append(caseArr, cases[j].CaseId)
		}
	}

	// tempCaseClass := defs.CaseClass{
	// 		Message: utils.TextCleaner(messages[i]), 
	// 		CaseIds: caseArr,
	// }
	// fmt.Println(tempCaseClass)
	// fmt.Println(messages[i])
	// classifiedMessagesArr = append(classifiedMessagesArr, messageClassifier(cases, messages[i]))

	return caseArr
}


func Tester(){
	var caseClassArr []defs.CaseClass
	// var query string
	// var caseArr []int
	cases := utils.ReadCases()
	messages := utils.ReadMessages()
	for i := range messages {
		tempCaseClass := defs.CaseClass{
				Message: utils.TextCleaner(messages[i]), 
				CaseIds: MessageClassifier(messages[i], cases),
		}

		caseClassArr = append(caseClassArr, tempCaseClass)
		// fmt.Println(tempCaseClass)
		// fmt.Println(messages[i])
		// classifiedMessagesArr = append(classifiedMessagesArr, messageClassifier(cases, messages[i]))
	}

	jsonString, _ := json.Marshal(caseClassArr)
	// log.Println(string(j))
	ioutil.WriteFile("./storage/output.json", jsonString, os.ModePerm)
}