package types

type Error struct {
	Message string `json:"message"`
}

type Response struct {
	Body any `json:"body"`
	Code int `json:"code"`
}
