package routes

type AdminGroup struct {
	GetUserInfo string
	GetAllUsers string

	AddRoleForUser    string
	RemoveRoleForUser string
}

var (
	Admin = AdminGroup{
		GetUserInfo: "/users/:user_uuid",
		GetAllUsers: "/users",

		AddRoleForUser:    "/users/:user_uuid/role/:role_name",
		RemoveRoleForUser: "/users/:user_uuid/role/:role_name",
	}
)
