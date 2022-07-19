# Publicsonar

Technical assignment for publicsonar comp
## Prepare 
```
    go mod init ehe.com/publicsonar
    go get github.com/gorilla/mux
    go get gopkg.in/Knetic/govaluate.v2
    go get github.com/tmdvs/Go-Emoji-Utils
```

---

## Usage 

1. run the local dev server : `go run main.go`
2. there exported postman collection json file : `publicsonar.postman_collection.json`
3. open json colletion file and send requests

---

## Endpoints

**/api/message-classifier**

> You can send a single to this endpoint and classify its case
- Method : `GET`
- Request Body : 
  - `"Message": "@praiseakinlami FIFA 23 dey come \nMan City vs PSG\nLet's see who people will choose"`
- Response Data :
  - `"CaseIds": [2]`
  - `"Message": "praiseakinlami fifa 23 dey come man city vs psg let s see who people will choose"`


**/api/messages-json**

> All messages in messages.json file can be classified by triggering this end point.
- Method : `GET`
- Request Body : `null`
- Response Data : `{"output.json file created ad ./storage directory !"}`

