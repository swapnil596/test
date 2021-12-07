package models

import "database/sql"

type (
	ShowUser struct {
		Id            string         `json:"id" form:"id"`
		Project_id    int            `json:"project_id" form:"project_id"`
		Name          string         `json:"name" form:"name"`
		Version       string         `json:"version" form:"version"`
		Url           sql.NullString `json:"url" form:"url"`
		Method        sql.NullString `json:"method" form:"method"`
		Protocol      string         `json:"protocol" form:"protocol"`
		Headers       sql.NullString `json:"headersy" form:"headers"`
		Request       sql.NullString `json:"request" form:"request"`
		Response      sql.NullString `json:"response" form:"response"`
		Degree        int            `json:"degree" form:"degree"`
		Active        bool           `json:"active" form:"active"`
		Created_by    string         `json:"created_by" form:"created_by"`
		Created_date  string         `json:"created_date" form:"created_date"`
		Modified_by   sql.NullString `json:"modified_by" form:"modified_by"`
		Modified_date sql.NullString `json:"modified_date" form:"modified_date"`
	}

	ReqUser struct {
		Id            string         `json:"id" form:"id"`
		Project_id    int            `json:"project_id" form:"project_id"`
		Name          string         `json:"name" form:"name"`
		Version       string         `json:"version" form:"version"`
		Url           sql.NullString `json:"url" form:"url"`
		Method        sql.NullString `json:"method" form:"method"`
		Protocol      string         `json:"protocol" form:"protocol"`
		Headers       sql.NullString `json:"headers" form:"headers"`
		Request       sql.NullString `json:"request" form:"request"`
		Response      sql.NullString `json:"response" form:"response"`
		Degree        int            `json:"degree" form:"degree"`
		Active        bool           `json:"active" form:"active"`
		Created_by    string         `json:"created_by" form:"created_by"`
		Created_date  string         `json:"created_date" form:"created_date"`
		Modified_by   sql.NullString `json:"modified_by" form:"modified_by"`
		Modified_date sql.NullString `json:"modified_date" form:"modified_date"`
	}
)
