package model

// TODO- keep in mind multitenancy and check if username or email (LDAP, SSO) should be the identifier

//TODO- Check model soft deleted vs delete (gorm.Models)

// User struct, which represents a user.
type User struct {
	Username string `gorm:"primarykey;size:100;unique;not null"`
	Surname  string `gorm:"size:100;not null"`
	Email    string `gorm:"size:255;unique;not null"`
}

// CreateUserRequest struct, which represents a user creation request.
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
