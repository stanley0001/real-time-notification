package services

import (
	"encoding/json"
	"log"
	"notification-service/models"
)

func HandleEvents(event string, data []byte) {
	switch event {
	case "create-comment":
		HandleCreateComment(data)
		break
	case "create-post":
		HandleCreatePost(data)
		break
	case "create-reaction":
		HandleCreateReaction(data)
		break
	case "update-post":
		HandleUpdatePost(data)
		break
	case "send-message":
		HandleSendMessage(data)
		break
	default:
		log.Printf("No implementation for event: %s", event)
	}
}

func HandleCreateComment(data []byte) {
	var comment models.Comments
	body := json.Unmarshal(data, &comment)
	log.Printf("notification for comment: %s", body)
	//handle comment
}
func HandleCreatePost(data []byte) {
	log.Printf("notification for post: %s", data)
	//handle post
}
func HandleUpdatePost(data []byte) {
	log.Printf("notification for post: %s", data)
	//handle post
}
func HandleCreateReaction(data []byte) {
	log.Printf("notification for reaction: %s", data)
	//handle reaction
}
func HandleSendMessage(data []byte) {
	log.Printf("notification for message: %s", data)
	//handle message
}
