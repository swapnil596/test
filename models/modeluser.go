package models

type (
	ShowUser struct {
		Id            int    `json:"id" form:"id"`
		Project_id    int    `json:"project_id,omitempty" form:"project_id"`
		Name          string `json:"name" form:"name"`
		Version       string `json:"version" form:"version"`
		Url           string `json:"url,omitempty" form:"url"`
		Method        string `json:"method,omitempty" form:"method"`
		Protocol      string `json:"protocol" form:"protocol"`
		Headers       string `json:"headers,omitempty" form:"headers"`
		Request       string `json:"request,omitempty" form:"request"`
		Response      string `json:"response,omitempty" form:"response"`
		Degree        int    `json:"degree" form:"degree"`
		Created_by    string `json:"created_by,omitempty" form:"created_by"`
		Created_date  string `json:"created_date,omitempty" form:"created_date"`
		Modified_by   string `json:"modified_by" form:"modified_by"`
		Modified_date string `json:"modified_date" form:"modified_date"`
	}

	ReqUser struct {
		Id            int    `json:"id" form:"id"`
		Project_id    int    `json:"project_id" form:"project_id"`
		Name          string `json:"name" form:"name"`
		Version       string `json:"version" form:"version"`
		Url           string `json:"url" form:"url"`
		Method        string `json:"method" form:"method"`
		Protocol      string `json:"protocol" form:"protocol"`
		Headers       string `json:"headers" form:"headers"`
		Request       string `json:"request" form:"request"`
		Response      string `json:"response" form:"response"`
		Degree        int    `json:"degree" form:"degree"`
		Created_by    string `json:"created_by" form:"created_by"`
		Created_date  string `json:"created_date" form:"created_date"`
		Modified_by   string `json:"modified_by" form:"modified_by"`
		Modified_date string `json:"modified_date" form:"modified_date"`
	}
)
