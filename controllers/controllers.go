package controllers

import(
	// "fmt"
    "io/ioutil"
	"os"
	"net/http"
	"encoding/json"
	// "github.com/gorilla/mux"
	"ehe.com/publicsonar/classifier"
	"ehe.com/publicsonar/utils"
	"ehe.com/publicsonar/defs"
)

func Classifier(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	cases := utils.ReadCases()
	var reqMessage defs.RequestMessage
	err := json.NewDecoder(r.Body).Decode(&reqMessage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	message := reqMessage.Message
	messageCases := classifier.MessageClassifier(message, cases)

	tempCaseClass := defs.CaseClass{
			Message: utils.TextCleaner(message), 
			CaseIds: messageCases,
	}
	res, _ := json.Marshal(tempCaseClass)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func MessagesJson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var caseClassArr []defs.CaseClass
	cases := utils.ReadCases()
	messages := utils.ReadMessages()
	for i := range messages {
		tempCaseClass := defs.CaseClass{
				Message: utils.TextCleaner(messages[i]), 
				CaseIds: classifier.MessageClassifier(messages[i], cases),
		}

		caseClassArr = append(caseClassArr, tempCaseClass)
	}
	jsonString, _ := json.Marshal(caseClassArr)
	ioutil.WriteFile("./storage/output.json", jsonString, os.ModePerm)

	res, _ := json.Marshal("output.json file created ad ./storage directory !")
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}