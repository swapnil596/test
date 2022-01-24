package models

import "database/sql"

type (
	ApiRegistration struct {
		Id            string         `json:"id"`
		ProjectId     string         `json:"project_id"`
		Name          string         `json:"name" form:"name"`
		Version       string         `json:"version" form:"version"`
		Url           sql.NullString `json:"url" form:"url"`
		Method        sql.NullString `json:"method" form:"method"`
		Protocol      string         `json:"protocol" form:"protocol"`
		Headers       sql.NullString `json:"headersy" form:"headers"`
		Request       sql.NullString `json:"requestBody" form:"requestBody"`
		Response      sql.NullString `json:"responseBody" form:"responseBody"`
		QueryParams   sql.NullString `json:"queryParameter" form:"queryParameter"`
		StatusCode    sql.NullInt64  `json:"status_code" form:"status_code"`
		TykUri        sql.NullString `json:"tykuri"`
		CacheTimeout  sql.NullString `json:"cacheTimeout"`
		RateLimit     sql.NullString `json:"rateLimit"`
		RateLimitPer  sql.NullString `json:"rateLimitPer"`
		Interval      sql.NullString `json:"interval"`
		Retries       sql.NullString `json:"retries"`
		Url2          sql.NullString `json:"url2"`
		AuthKey       sql.NullString `json:"authkey"`
		CacheByHeader bool           `json:"cache_by_header"`
		AuthType      int            `json:"authtype" form:"authtype"`
		Username      sql.NullString `json:"username"`
		Password      sql.NullString `json:"password"`
		Degree        int            `json:"degree" form:"degree"`
		CreatedBy     string         `json:"created_by" form:"created_by"`
		CreatedDate   string         `json:"created_date" form:"created_date"`
		ModifiedBy    sql.NullString `json:"modified_by" form:"modified_by"`
		ModifiedDate  sql.NullString `json:"modified_date" form:"modified_date"`
		Active        bool           `json:"active" form:"active"`
	}

	TempApi struct {
		Id            string                 `json:"id" form:"id"`
		ProjectId     string                 `json:"project_id" form:"project_id"`
		Name          string                 `json:"name" form:"name"`
		Version       string                 `json:"version" form:"version"`
		Url           string                 `json:"url" form:"url"`
		Method        string                 `json:"method" form:"method"`
		Protocol      string                 `json:"protocol" form:"protocol"`
		Headers       map[string]interface{} `json:"headers" form:"headers"`
		Request       map[string]interface{} `json:"requestBody" form:"requestBody"`
		Response      map[string]interface{} `json:"responseBody" form:"responseBody"`
		QueryParams   map[string]interface{} `json:"queryParameter" form:"queryParameter"`
		TykUri        sql.NullString         `json:"tykuri"`
		CacheTimeout  string                 `json:"cacheTimeout"`
		RateLimit     string                 `json:"rateLimit"`
		RateLimitPer  string                 `json:"rateLimitPer"`
		Interval      string                 `json:"interval"`
		Retries       string                 `json:"retries"`
		Url2          string                 `json:"url2"`
		AuthKey       string                 `json:"authkey"`
		CacheByHeader bool                   `json:"cache_by_header"`
		AuthType      int                    `json:"authtype" form:"authtype"`
		Username      string                 `json:"username"`
		Password      string                 `json:"password"`
		Degree        int                    `json:"degree" form:"degree"`
		Active        bool                   `json:"active" form:"active"`
		CreatedBy     string                 `json:"created_by" form:"created_by"`
		CreatedDate   string                 `json:"created_date" form:"created_date"`
		ModifiedBy    sql.NullString         `json:"modified_by" form:"modified_by"`
		ModifiedDate  sql.NullString         `json:"modified_date" form:"modified_date"`
	}
)
