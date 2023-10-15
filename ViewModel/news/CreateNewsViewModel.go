package news

type CreateNewsViewModel struct {
	Title            string `form:"Title" validate:"required"`
	ShortDescription string `form:"ShortDescription" validate:"required"`
	Description      string `form:"Description" validate:"required"`
	ImageName        string
	CreatorUserId    string
}

type EditNewsViewModel struct {
	Id               string
	Title            string `form:"Title" validate:"required"`
	ShortDescription string `form:"ShortDescription" validate:"required"`
	Description      string `form:"Description" validate:"required"`
	ImageName        string
	CreatorUserId    string
}
