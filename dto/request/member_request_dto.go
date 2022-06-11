package dto

type MemberCreate struct {
	Email    string `json:"email" validate:"lte=50,required"`
	Password string `json:"password" validate:"lte=100,required"`
	Name     string `json:"name" validate:"lte=50,required"`
	Mobile   string `json:"mobile" validate:"required"`
}

type SearchMemberQueryParams struct {
}
