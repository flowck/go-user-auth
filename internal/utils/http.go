package utils

import "encoding/json"

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) []byte {
	response, err := json.Marshal(MessageResponse{Message: message})

	if err != nil {
		panic(err)
	}

	return response
}
