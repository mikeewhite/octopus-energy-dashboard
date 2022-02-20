package octopusenergy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const baseURL = "https://api.octopus.energy"

// Client for interfacing with Octopus Energy API
type Client struct {
	apiKey                       string
	electricityMeterMPAN         string
	electricityMeterSerialNumber string
	httpClient                   *http.Client
}

// New creates a new client
func New(apiKey, electricityMeterMPAN, electricityMeterSerialNumber string) *Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &Client{
		apiKey:                       apiKey,
		electricityMeterMPAN:         electricityMeterMPAN,
		electricityMeterSerialNumber: electricityMeterSerialNumber,
		httpClient:                   httpClient,
	}
}

// GetElectricityMeterConsumption TODO
func (c *Client) GetElectricityMeterConsumption() (*Consumption, error) {
	url := fmt.Sprintf("%s/v1/electricity-meter-points/%s/meters/%s/consumption",
		baseURL,
		c.electricityMeterMPAN,
		c.electricityMeterSerialNumber)
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response (%s) received on retrieving electricity meter consumption", resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error on reading request body")
	}

	var consumption Consumption
	err = json.Unmarshal(body, &consumption)
	if err != nil {
		return nil, err
	}

	return &consumption, nil
}
