package models

// Session ...
type Session struct {
	ID       int    `json:"id"`
	GroupID  int    `json:"group_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Admin    *bool  `json:"admin"`
}

// IsAdmin ..
func (s *Session) IsAdmin() bool {
	return *s.Admin
}
