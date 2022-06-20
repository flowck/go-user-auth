package user

type User struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
	Roles     []Role
}

type Role struct {
	Name string
}
