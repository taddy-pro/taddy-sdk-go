package taddy

func (client *Client) GetExchangeFeed(user *User, options *GetExchangeRequestOptions) ([]*ExchangeItem, error) {
	var res *GetExchangeResponse
	if options == nil {
		options = &GetExchangeRequestOptions{}
	}
	err := client.call(POST, "/exchange/feed", &GetExchangeRequest{
		BasicRequest: BasicRequest{
			PubId: client.pubId,
			User:  *user,
		},
		GetExchangeRequestOptions: *options,
	}, &res)
	return res.Result, err
}

func (client *Client) SendExchangeImpressions(ids []string) error {
	return client.call(POST, "/exchange/impressions", &SendExchangeImpressionRequest{
		Ids: ids,
	}, nil)
}

type GetExchangeRequest struct {
	BasicRequest
	GetExchangeRequestOptions
}

type ExchangeItem struct {
	Id          string   `json:"id"`
	Uid         string   `json:"uid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       *string  `json:"image"`
	FullImage   *string  `json:"fullImage"`
	Type        string   `json:"type"`
	Price       *float64 `json:"price"`
	Link        string   `json:"link"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	ExpiresAt   string   `json:"expiresAt"`
}

type GetExchangeRequestOptions struct {
	Limit           int    `json:"limit,default=4"`
	ImageFormat     string `json:"imageFormat,default=webp"`
	AutoImpressions bool   `json:"autoImpressions,default=false"`
}

type GetExchangeResponse struct {
	Result []*ExchangeItem `json:"result"`
}

type SendExchangeImpressionRequest struct {
	Ids []string `json:"ids"`
}
