package controllers

import(
	// "fmt"
	"net/http"
	"encoding/json"
	// "github.com/gorilla/mux"
	"ehe.com/publicsonar/classifier"
	"ehe.com/publicsonar/utils"
	"ehe.com/publicsonar/defs"
)

func Classifier(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var reqMessage defs.RequestMessage
	err := json.NewDecoder(r.Body).Decode(&reqMessage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	message := reqMessage.Message
	cases := classifier.MessageClassifier(message)

	tempCaseClass := defs.CaseClass{
			Message: utils.TextCleaner(message), 
			CaseIds: cases,
	}
	res, _ := json.Marshal(tempCaseClass)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}