package models

type FlowxpertAPIMaster struct {
	Id           string `dynamo:"id"`
	Name         string `dynamo:"name"`
	Version      string `dynamo:"version"`
	Url          string `dynamo:"url"`
	Headers      string `dynamo:"headers"`
	Method       string `dynamo:"method"`
	Params       string `dynamo:"params"`
	RequestBody  string `dynamo:"request_body"`
	ResponseBody string `dynamo:"response_body"`
	APIType      string `dynamo:"api_type"`
	Active       bool   `dynamo:"active"`
	Published    bool   `dynamo:"published"`
	ProjectId    int    `dynamo:"project_id"`
	TypeOf       string `dynamo:"type"`
	CreatedBy    string `dynamo:"created_by"`
	CreatedDate  string `dynamo:"created_date"`
	ModifiedBy   string `dynamo:"modified_by"`
	ModifiedDate string `dynamd:"modified_date"`
}
