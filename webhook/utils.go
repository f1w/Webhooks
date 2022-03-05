package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
		if res.StatusCode == http.StatusTooManyRequests {
			rateLimitPayload := struct{
				RetryAfter float64 `json:"retry_after"`
			}{}
			time.Sleep(time.Duration(int(math.Ceil(rateLimitPayload.RetryAfter))) * time.Second)
			return s.MakeRequest(method, endpoint, Body)
		}
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
