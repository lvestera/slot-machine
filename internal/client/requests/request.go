package requests

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/lvestera/slot-machine/internal/models"
)

const getConfigUrl = "/get-config"
const sendResultUrl = "/send-result"

type RequestClient struct {
	Host   string
	Client resty.Client
}

func GetRequestClient(host string) *RequestClient {
	client := resty.New()
	client.SetBaseURL(host)
	client.SetHeader("Content-Type", "application/json")

	return &RequestClient{
		Host:   host,
		Client: *client,
	}
}

func (rc *RequestClient) GetConfig() ([]models.Coefficient, error) {
	log.Println("Request config")

	resp, err := rc.Client.NewRequest().Post(getConfigUrl)

	if err != nil {
		return nil, err
	}

	var coefficients []models.Coefficient

	err = json.Unmarshal(resp.Body(), &coefficients)
	if err != nil {
		return nil, err
	}

	return coefficients, nil
}

func (rc *RequestClient) SaveResults(results map[int]models.Result) error {

	var err error
	var body []byte

	if body, err = json.Marshal(results); err != nil {
		return err
	}

	_, err = rc.Client.NewRequest().SetBody(body).Post(sendResultUrl)
	if err != nil {
		return err
	}

	return nil
}
