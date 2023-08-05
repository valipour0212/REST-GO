package user

type CreateNewUserViewModel struct {
	FirstName     string
	LastName      string `validate:"required"`
	Email         string
	UserName      string
	Password      string
	CreatorUserId string
}
