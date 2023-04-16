package session

type SessionRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type SessionResponse struct {
	Token string `json:"token"`
}

type SessionRecordSet struct {
	Id string
}
