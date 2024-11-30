package constants

type UserRole string

const (
	AdminUserRole    UserRole = "ADMIN"
	EmployeeUserRole UserRole = "EMPLOYEE"
	CustomerUserRole UserRole = "CUSTOMER"
)

type Scope string

const (
	UsersReadScope   Scope = "users:read"
	UsersUpdateScope Scope = "users:update"
	UsersCreateScope Scope = "users:create"
	UsersDeleteScope Scope = "users:delete"
)
