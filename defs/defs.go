package defs

type Cases struct{
    CaseId    int `json:"case_id"`
    Query string `json:"query"`
    Queries string `json:"queries"`
}

type Exp struct{
	Order int
	Value string
}

type CaseClass struct{
	CaseIds []int
	Message string
}

type RequestMessage struct{
	Message string
}