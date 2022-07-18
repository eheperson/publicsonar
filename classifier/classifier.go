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
)


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
	message = utils.TextCleaner(message)
	tokenizedQuery := QueryTokenizer(query)
	for _, val := range tokenizedQuery{
		if strings.Contains(message, val){
			query = strings.Replace(query, val, "true" , 1)
		} else {
			query = strings.Replace(query, val, "false" , 1)
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

func MessageClassifier(message string, ) []int{
	var query string
	var caseArr []int
	cases := utils.ReadCases()
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

			refactoredQuery := QueryRefactor(query, messages[i])
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