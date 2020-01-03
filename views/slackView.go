package views

// NewSlack ...
type NewSlack struct {
	FulfillmentText string  `json:"fulfillmentText,omitempty"`
	Payload         Payload `json:"payload,omitempty"`
}

// Payload ...
type Payload struct {
	Slack Slack `json:"slack,omitempty"`
}

// Slack ...
type Slack struct {
	Text        string        `json:"text,omitempty"`
	Attachments []Attachments `json:"attachments,omitempty"`
}

// Attachments ...
type Attachments struct {
	Text           string    `json:"text,omitempty"`
	Title          string    `json:"title,omitempty"`
	Fallback       string    `json:"fallback,omitempty"`
	CallbackID     string    `json:"callback_id,omitempty"`
	Color          string    `json:"color,omitempty"`
	AttachmentType string    `json:"attachment_type,omitempty"`
	Actions        []Actions `json:"actions,omitempty"`
	ThumbURL       string    `json:"thumb_url,omitempty"`
	Footer         string    `json:"footer,omitempty"`
}

// Actions ...
type Actions struct {
	Name  string `json:"name,omitempty"`
	Text  string `json:"text,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Style string `json:"style,omitempty"`
}
