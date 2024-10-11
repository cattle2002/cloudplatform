package protocol

type LoginRetMsg struct {
	LoginCode string `json:"LoginCode"`
	Secret    string `json:"Secret"`
}
