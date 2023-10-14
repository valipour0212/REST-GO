package user

type CreateNewUserViewModel struct {
	FirstName     string
	LastName      string `validate:"required"`
	Email         string
	UserName      string
	Password      string
	CreatorUserId string
}

type EditUserViewModel struct {
	TargetUserID string
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Email        string `validate:"required"`
	UserName     string `validate:"required"`
	Password     string `validate:"required"`
}

type EditUserRoleViewModel struct {
	TargetUserId string
	Roles        []string `validate:"required"`
}

type EditUserPasswordViewModel struct {
	TargetUserId string
	Password     string `validate:"required"`
}
