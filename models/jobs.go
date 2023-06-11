package models

type CTSJobs struct {
	JobTitle    string `json:"job_title"`
	CompanyName string `json:"company_name"`
	JobID       string `json:"job_id"`
	Description string `json:"description"`
	PostedTime  string `json:"posted_at"`
	JobURL      string `json:"job_url"`
}
