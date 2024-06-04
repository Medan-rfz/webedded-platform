package users

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "moderator"
	Teacher   Role = "teacher"
	Student   Role = "student"
)
