package task

type DTO struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Created_date   string `json:"created_date"`
	Completed_date string `json:"completed_date"`
	Deleted_date   string `json:"deleted_date"`
}

type Form struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
