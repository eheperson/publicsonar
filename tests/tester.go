package main

import (
    "encoding/json"
    // "fmt"
    "io/ioutil"
    "log"
	"os"
	// "strings"
	// "regexp"
	"gopkg.in/Knetic/govaluate.v2"
	"ehe.com/publicsonar/utils"
	"ehe.com/publicsonar/defs"
	"ehe.com/publicsonar/classifier"
)


func main() {
	var caseClassArr []defs.CaseClass
	var query string
	var caseArr []int
	cases := utils.ReadCases()
	messages := utils.ReadMessages()
	for i := range messages {
		caseArr = nil
		for j := range cases {
			// fmt.Println(cases[j])
			if len(cases[j].Query) != 0{
				query = cases[j].Query
			}
			if len(cases[j].Queries) != 0{
				query = cases[j].Queries
			}

			refactoredQuery := classifier.QueryRefactor(query, messages[i])
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

		tempCaseClass := defs.CaseClass{
				Message: utils.TextCleaner(messages[i]), 
				CaseIds: caseArr,
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