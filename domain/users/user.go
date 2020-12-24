package users

// User details
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	DataCreated string `json:"datecreated"`
}
