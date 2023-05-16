package aboutsh

type CreateReq struct {
	Title      string `json:"title"`
	Intro      string `json:"intro"`
	LinkedIn   string `json:"linkedin"`
	Youtube    string `json:"youtube"`
	Facebook   string `json:"facebook"`
	Telegram   string `json:"telegram"`
	ResumeLink string `json:"resumelink"`
}

type UpdateReq struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Intro      string `json:"intro"`
	LinkedIn   string `json:"linkedin"`
	Youtube    string `json:"youtube"`
	Facebook   string `json:"facebook"`
	Telegram   string `json:"telegram"`
	ResumeLink string `json:"resumelink"`
}

type DeleteReq struct {
	Id int `json:"id"`
}

type FullResponseBody struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Intro      string `json:"intro"`
	LinkedIn   string `json:"linkedin"`
	Youtube    string `json:"youtube"`
	Facebook   string `json:"facebook"`
	Telegram   string `json:"telegram"`
	ResumeLink string `json:"resumelink"`
}

type FullResponse struct {
	ErrorCode    int               `json:"error_code"`
	ErrorMessage string            `json:"error_message"`
	Body         *FullResponseBody `json:"body"`
}
