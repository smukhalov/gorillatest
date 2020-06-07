package frombackend

type Error struct {
	Message []string `json:"message"`
}

type Errors struct {
	Error `json:"errors"`
}
