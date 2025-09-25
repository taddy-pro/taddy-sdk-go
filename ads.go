package taddy

func (client *Client) GetAd(user *User, options *GetAdRequestOptions) (*Ad, error) {
	var res *GetAdResponse
	if options == nil {
		options = &GetAdRequestOptions{
			Format: "bot-ad",
		}
	}
	err := client.call(POST, "/ads/get", &GetAdRequest{
		BasicRequest: BasicRequest{
			PubId: client.pubId,
			User:  *user,
		},
		GetAdRequestOptions: *options,
	}, &res)
	return res.Result, err
}

func (client *Client) SendAdImpression(id string) error {
	return client.call(POST, "/ads/impressions", &SendAdImpressionRequest{
		Id: id,
	}, nil)
}

type Ad struct {
	Id          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Video       *string `json:"video"`
	Icon        *string `json:"icon"`
	Text        *string `json:"text"`
	Button      *string `json:"button"`
	Link        string  `json:"link"`
}

type GetAdRequest struct {
	BasicRequest
	GetAdRequestOptions
}

type GetAdRequestOptions struct {
	Format string `json:"format"`
}

type GetAdResponse struct {
	Result *Ad `json:"result"`
}

type SendAdImpressionRequest struct {
	Id string `json:"id"`
}
