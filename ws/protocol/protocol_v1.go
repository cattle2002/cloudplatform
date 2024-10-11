package protocol

type LoginMsg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
