package giant_bomb

type BaseResponse struct {
	Error                string `json:"error"`
	Limit                int    `json:"limit"`
	Offset               int    `json:"offset"`
	NumberOfPageResults  int    `json:"number_of_page_results"`
	NumberOfTotalResults int    `json:"number_of_total_results"`
	StatusCode           int    `json:"status_code"`
	Version              string `json:"version"`
}

type GameRatings struct {
	BaseResponse
	Results []struct {
		APIDetailURL string      `json:"api_detail_url"`
		ID           int         `json:"id"`
		GUID         string      `json:"guid"`
		Image        interface{} `json:"image"`
		Name         string      `json:"name"`
		RatingBoard  struct {
			APIDetailURL string `json:"api_detail_url"`
			ID           int    `json:"id"`
			Name         string `json:"name"`
		} `json:"rating_board"`
	} `json:"results"`
}
