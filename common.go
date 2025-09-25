package taddy

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

type User struct {
	Id        int64   `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Username  *string `json:"username,omitempty"`
	Premium   bool    `json:"premium,omitempty"`
	Language  string  `json:"language,omitempty"`
	Country   *string `json:"country,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	IP        *string `json:"ip,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	BirthDate *string `json:"birthDate,omitempty"`
}

type BasicRequest struct {
	PubId string `json:"pubId"`
	User  User   `json:"user"`
}

type StartRequest struct {
	BasicRequest
	Start string `json:"start"`
}

type GetInitDataRequest struct {
	BasicRequest
}

type ResourceInitData struct {
	Id            int64    `json:"id"`
	Username      string   `json:"username"`
	Apps          []string `json:"apps"`
	ExternalAds   bool     `json:"externalAds"`
	TeleAdsToken  string   `json:"teleAdsToken"`
	TeleAdsUnitId string   `json:"teleAdsUnitId"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (client *Client) Start(user *User, start string) error {
	return client.call(POST, "/start", &StartRequest{
		BasicRequest: BasicRequest{
			PubId: client.pubId,
			User:  *user,
		},
		Start: start,
	}, nil)
}

func (client *Client) GetInitData(user *User) (*ResourceInitData, error) {
	var result *ResourceInitData
	err := client.call(POST, "/start", &GetInitDataRequest{
		BasicRequest: BasicRequest{
			PubId: client.pubId,
			User:  *user,
		},
	}, &result)
	return result, err
}
