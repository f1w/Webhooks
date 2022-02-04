package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DiscordApiVersion = "v9"
	DiscordBaseUrl  ="https://discord.com/api/"
)



func (s *Session) MakeRequest(method string, endpoint string, Body []byte) (*http.Response, error) {
	var body io.Reader

	if Body == nil {
		body = nil
	} else {
		body = bytes.NewBuffer(Body)
	}

	reqUrl := fmt.Sprintf("%s%s%s", DiscordBaseUrl, DiscordApiVersion, endpoint)
	req, err := http.NewRequest(method, reqUrl, body)


	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %v",s.Token))

	res, err := s.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errResp, readErr := ioutil.ReadAll(res.Body)

		if readErr != nil {
			return nil, readErr
		}

		return nil, errors.New(fmt.Sprintf("Invalid Request - Status Code: %d - Error: %s", res.StatusCode, string(errResp)))
	}

	return res, nil
}

func ParseJsonFromHttpResponse(response *http.Response, out interface{}) error {

	err := json.NewDecoder(response.Body).Decode(&out)

	if err != nil {
		return err
	}
	return nil
}

func (s *Session) ExecuteLogWebhook(webhookId Snowflake, webhookName string, logType *LogType) error {

	webhook, err := s.GetWebhook(webhookId)

	if err != nil {
		return err
	}

	err = s.ExecuteWebhook(webhook.Id, webhook.Token,
		&WebhookMessage{
			Username: webhookName,
			Embeds: &[]Embed{
				{
					Title:       logType.LogTitle,
					Description: logType.LogDescription,
					Fields: logType.LogFields,
					Timestamp: time.Now().UTC(),
				},
			},
		})

	if err != nil {
		return err
	}

	return nil
}

