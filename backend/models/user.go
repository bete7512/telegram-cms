package models

type User struct {
	Id        string `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"not null" `
	LastName  string `json:"last_name" gorm:"not null" `
	Email     string `json:"email" gorm:"unique;not null" `
	Password  string `json:"password" gorm:"not null"`
	Status    string `json:"status" gorm:"not null; default:'false'"`
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
