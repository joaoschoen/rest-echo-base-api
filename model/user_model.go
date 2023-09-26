package model

// TYPES
type UnsafeUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SafeUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// RESPONSE OBJECTS
type GetUserResponse struct {
	Data SafeUser
}

type GetUserListResponse struct {
	Data   []SafeUser
	Paging Paging
}

type PostUserResponse struct {
	ID string
}

type PutUserResponse struct {
	Data SafeUser
}

type DeleteUserResponse struct {
	ID string
}
