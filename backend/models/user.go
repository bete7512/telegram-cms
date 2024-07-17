package models

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Role struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UserRoles struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type Permission struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RolePermissions struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}
