package constants

type RoleName string

const (
	BASIC RoleName = "BASIC"
	ADMIN RoleName = "ADMIN"
)

var ValidRoleNames = []RoleName{
	BASIC,
	ADMIN,
}