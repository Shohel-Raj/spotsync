package httpresponse

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Details string `json:"details,omitempty"`
}
