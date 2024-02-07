package model

type BaseResponse struct {
	Error                string `json:"error"`
	Limit                int    `json:"limit"`
	Offset               int    `json:"offset"`
	NumberOfPageResults  int    `json:"number_of_page_results"`
	NumberOfTotalResults int    `json:"number_of_total_results"`
	StatusCode           int    `json:"status_code"`
	Version              string `json:"version"`
}

type Game struct {
	Aliases                interface{} `json:"aliases"`
	APIDetailURL           string      `json:"api_detail_url"`
	DateAdded              string      `json:"date_added"`
	DateLastUpdated        string      `json:"date_last_updated"`
	Deck                   interface{} `json:"deck"`
	Description            interface{} `json:"description"`
	ExpectedReleaseDay     interface{} `json:"expected_release_day"`
	ExpectedReleaseMonth   interface{} `json:"expected_release_month"`
	ExpectedReleaseQuarter interface{} `json:"expected_release_quarter"`
	ExpectedReleaseYear    interface{} `json:"expected_release_year"`
	GUID                   string      `json:"guid"`
	ID                     int         `json:"id"`
	Image                  struct {
		IconURL        string `json:"icon_url"`
		MediumURL      string `json:"medium_url"`
		ScreenURL      string `json:"screen_url"`
		ScreenLargeURL string `json:"screen_large_url"`
		SmallURL       string `json:"small_url"`
		SuperURL       string `json:"super_url"`
		ThumbURL       string `json:"thumb_url"`
		TinyURL        string `json:"tiny_url"`
		OriginalURL    string `json:"original_url"`
		ImageTags      string `json:"image_tags"`
	} `json:"image"`
	ImageTags []struct {
		APIDetailURL string `json:"api_detail_url"`
		Name         string `json:"name"`
		Total        int    `json:"total"`
	} `json:"image_tags"`
	Name                string      `json:"name"`
	NumberOfUserReviews int         `json:"number_of_user_reviews"`
	OriginalGameRating  interface{} `json:"original_game_rating"`
	OriginalReleaseDate interface{} `json:"original_release_date"`
	Platforms           []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
		Abbreviation  string `json:"abbreviation"`
	} `json:"platforms"`
	SiteDetailURL string `json:"site_detail_url"`
	Images        []struct {
		IconURL   string `json:"icon_url"`
		MediumURL string `json:"medium_url"`
		ScreenURL string `json:"screen_url"`
		SmallURL  string `json:"small_url"`
		SuperURL  string `json:"super_url"`
		ThumbURL  string `json:"thumb_url"`
		TinyURL   string `json:"tiny_url"`
		Original  string `json:"original"`
		Tags      string `json:"tags"`
	} `json:"images"`
	Videos     []interface{} `json:"videos"`
	Characters interface{}   `json:"characters"`
	Concepts   []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"concepts"`
	Developers []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"developers"`
	FirstAppearanceCharacters interface{} `json:"first_appearance_characters"`
	FirstAppearanceConcepts   []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"first_appearance_concepts"`
	FirstAppearanceLocations interface{} `json:"first_appearance_locations"`
	FirstAppearanceObjects   interface{} `json:"first_appearance_objects"`
	FirstAppearancePeople    interface{} `json:"first_appearance_people"`
	Franchises               interface{} `json:"franchises"`
	Genres                   []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"genres"`
	KilledCharacters interface{} `json:"killed_characters"`
	Locations        interface{} `json:"locations"`
	Objects          interface{} `json:"objects"`
	People           interface{} `json:"people"`
	Publishers       []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"publishers"`
	Releases []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"releases"`
	SimilarGames []struct {
		APIDetailURL  string `json:"api_detail_url"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SiteDetailURL string `json:"site_detail_url"`
	} `json:"similar_games"`
}

type Rating struct {
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
}

type AutoGenerated struct {
	BaseResponse
	Game struct {
		Game
	} `json:"results"`
}

type GameRatings struct {
	BaseResponse
	Ratings []struct {
		Rating
	} `json:"results"`
}
