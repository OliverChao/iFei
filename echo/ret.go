package echo

// structs

type RetResult struct {
	Code  int                    `json:"code"`
	Msg   string                 `json:"msg"`
	Data  interface{}            `json:"data"`
	Extra map[string]interface{} `json:"extra_data"`
}

func NewRetResult() (ret *RetResult) {
	return &RetResult{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

// be used for testing post article binding data
type BindPostArticle struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Tags        string `json:"tags" binding:"required"`
	Commentable bool   `json:"commentable" binding:"required"`
	Stared      bool   `json:"stared" binding:"required"`
	Topped      bool   `json:"topped" binding:"required"`
}

type BindLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
