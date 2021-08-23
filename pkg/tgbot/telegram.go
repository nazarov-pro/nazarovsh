package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	reqBody := &sendMessageReqBody{
		ChatID: body.Message.Chat.ID,
		Text:   body.Message.Text,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	apiKey := os.Getenv("TGBOT_API_KEY")

	// Send a post request with your token
	response, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiKey), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sayPolo(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// sayPolo takes a chatID and sends "polo" to them
func sayPolo(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Polo!!",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	apiKey := os.Getenv("TGBOT_API_KEY")

	// Send a post request with your token
	res, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", apiKey), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
