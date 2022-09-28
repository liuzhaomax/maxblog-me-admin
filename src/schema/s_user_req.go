package schema

type LoginInfo struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}
