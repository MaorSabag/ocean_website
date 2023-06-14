package models

type Repository struct {
	Name        string `json:"Name"`
	Languange   string `json:"Language"`
	Stars       int    `json:"Stars"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
	ReleaseDate string `json:"ReleaseDate"`
}

type Response struct {
	FulfillmentText     string `json:"fulfillmentText"`
	FulfillmentMessages []struct {
		Text struct {
			Text []string `json:"text"`
		} `json:"text"`
	} `json:"fulfillmentMessages"`
	Source  string `json:"source"`
	Payload struct {
		Google struct {
			ExpectUserResponse bool `json:"expectUserResponse"`
			RichResponse       struct {
				Items []struct {
					SimpleResponse struct {
						TextToSpeech string `json:"textToSpeech"`
					} `json:"simpleResponse"`
				} `json:"items"`
			} `json:"richResponse"`
		} `json:"google"`
		Facebook struct {
			Text string `json:"text"`
		} `json:"facebook"`
		Slack struct {
			Text string `json:"text"`
		} `json:"slack"`
	} `json:"payload"`
}

type Database []Repository
