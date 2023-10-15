package news

type CreateNewsViewModel struct {
	Title            string `form:"Title" validate:"required"`
	ShortDescription string `form:"ShortDescription" validate:"required"`
	Description      string `form:"Description" validate:"required"`
	ImageName        string
	CreatorUserId    string
}
