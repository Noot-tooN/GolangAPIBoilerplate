package routes

type AdminGroup struct {
	GetUserInfo string
	GetAllUsers string
}

var (
	Admin = AdminGroup{
		GetUserInfo: "/users/:user_id",
		GetAllUsers: "/users",
	}
)
