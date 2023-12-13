package routes

type UserGroup struct {
	Register   string
	Login string
	SoftDelete string
	HardDelete string
}

var (
	User = UserGroup{
		Register:   "/register",
		Login: "/login",
		SoftDelete: "/soft-delete",
		HardDelete: "/hard-delete",
	}
)