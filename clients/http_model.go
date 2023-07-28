package clients

const (
	CodeOK = 1
)

type Type int8

const (
	Others Type = iota
	Image
	Audio
	Video
	Document
)

type Resource struct {
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
	Type      Type   `json:"type"`
}

type Response struct {
	RID     string `json:"rid"`
	CID     string `json:"cid"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          string   `json:"id"`           //
	Token       string   `json:"token"`        //
	TokenExpire int      `json:"token_expire"` //
	Username    string   `json:"username"`     //
	Nickname    string   `json:"nickname"`     //
	Avatar      Resource `json:"avatar"`       //
	BIO         string   `json:"bio"`          //
}

type LoginResponseWrapper struct {
	Response
	Data LoginResponse `json:"data"`
}

type RegisterRequest struct {
	Username     string `json:"username"`      //
	Password     string `json:"password"`      //
	CaptchaType  int    `json:"captcha_type"`  //
	CaptchaID    string `json:"captcha_id"`    //
	CaptchaName  string `json:"captcha_name"`  //
	CaptchaValue string `json:"captcha_value"` //
}

type RegisterResponse struct {
	ID string `json:"id"`
}

type RegisterResponseWrapper struct {
	Response
	Data RegisterResponse `json:"data"`
}

//TODO

type OAuth2AuthorizeRequest struct{}
type OAuth2AuthorizeResponse struct{}

type OAuth2TokenRequest struct{}
type OAuth2TokenResponse struct{}
