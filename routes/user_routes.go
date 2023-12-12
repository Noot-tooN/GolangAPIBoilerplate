package routes

type UserGroup struct {
	Register   string
	Login string
}

var (
	User = UserGroup{
		Register:   "/register",
		Login: "/login",
	}
)