package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Session struct {
	httpClient *http.Client
	Token string
}

func Connect(token string) *Session {
	return &Session{
		httpClient: &http.Client{},
		Token: token,
	}
}


func (s *Session) CreateWebhook(name string, channelId Snowflake, avatar ...string) (*Webhook, error){

	var webhook Webhook

	bodyPayload := map[string]string{"name": name}

	jsonPayload, err := json.Marshal(bodyPayload)

	if err != nil {
		return &webhook, err
	}

	response, err := s.MakeRequest("POST", fmt.Sprintf("/channels/%v/webhooks", channelId), jsonPayload)

	if err != nil {

		return &webhook, err
	}

	err = ParseJsonFromHttpResponse(response, &webhook)

	return &webhook, nil

}


func (s *Session) GetWebhook(id Snowflake) (*Webhook, error) {

	var webhook Webhook

	response, err := s.MakeRequest("GET", fmt.Sprintf("/webhooks/%v", id), nil)

	if err != nil {
		return &webhook, err
	}

	err = ParseJsonFromHttpResponse(response, &webhook)

	return &webhook, nil

}


func (s *Session) DeleteWebhook(id Snowflake) (*Webhook, error) {

	var webhook Webhook

	response, err := s.MakeRequest("DELETE", fmt.Sprintf("/webhooks/%v", id), nil)

	if err != nil {
		return &webhook, err
	}

	err = ParseJsonFromHttpResponse(response, &webhook)

	return &webhook, nil

}

func (s *Session) ExecuteWebhook(id Snowflake, token string, messageHook *WebhookMessage) error {

	requestBody, err := json.Marshal(&messageHook)

	if err != nil {
		return err
	}

	_, reqErr := s.MakeRequest("POST", fmt.Sprintf("/webhooks/%v/%v", id, token), requestBody)

	if reqErr != nil {
		return  reqErr
	}

	return nil
}


func (s *Session) GetChannelWebhooks(channelId Snowflake) (*[]Webhook, error) {

	var webhooks []Webhook

	response, err := s.MakeRequest("GET", fmt.Sprintf("/channels/%v/webhooks", channelId), nil)

	if err != nil {
		return &webhooks, err
	}

	err = ParseJsonFromHttpResponse(response, &webhooks)

	return &webhooks, nil

}


func (s *Session) GetGuildWebhooks(guildId Snowflake) (*[]Webhook, error) {

	var webhooks []Webhook

	response, err := s.MakeRequest("GET", fmt.Sprintf("/guilds/%v/webhooks", guildId), nil)

	if err != nil {
		return &webhooks, err
	}

	err = ParseJsonFromHttpResponse(response, &webhooks)

	return &webhooks, nil

}

