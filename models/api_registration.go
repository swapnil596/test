package models

import "database/sql"

type (
	ApiRegistration struct {
		Id           string         `json:"id"`
		ProjectId    int            `json:"project_id"`
		Name         string         `json:"name" form:"name"`
		Version      string         `json:"version" form:"version"`
		Url          sql.NullString `json:"url" form:"url"`
		Method       sql.NullString `json:"method" form:"method"`
		Protocol     string         `json:"protocol" form:"protocol"`
		Headers      sql.NullString `json:"headersy" form:"headers"`
		Request      sql.NullString `json:"requestBody" form:"requestBody"`
		Response     sql.NullString `json:"responseBody" form:"responseBody"`
		QueryParams  sql.NullString `json:"queryParameter" form:"queryParameter"`
		StatusCode   sql.NullInt64  `json:"status_code" form:"status_code"`
		Degree       int            `json:"degree" form:"degree"`
		CreatedBy    string         `json:"created_by" form:"created_by"`
		CreatedDate  string         `json:"created_date" form:"created_date"`
		ModifiedBy   sql.NullString `json:"modified_by" form:"modified_by"`
		ModifiedDate sql.NullString `json:"modified_date" form:"modified_date"`
		Active       bool           `json:"active" form:"active"`
		RateLimit    sql.NullInt64  `json:"rate_limit" form:"rate_limit"`
	}
)
