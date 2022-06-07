package dto

type MemberCreate struct {
	Email    string `json:"code" validate:"lte=50,required"`
	Password string `json:"type" validate:"lte=100,required"`
	Name     string `json:"name" validate:"lte=50,required"`
	Mobile   string `json:"telNo" validate:"required,telNo"`
}

type SearchMemberQueryParams struct {
}
