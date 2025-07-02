package payment

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/common"
	"github.com/alielmi98/go-ecommerce-api/config"
)

// Zarinpal is a struct that implements the PaymentGateway interface

// Payment Request Response Dto
type PaymentRequestResponse struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Authority string   `json:"authority"`
	Errors    []string `json:"errors,omitempty"`
}

type Zarinpal struct {
	cfg *config.Config
}

func NewZarinpalGateway(cfg *config.Config) *Zarinpal {
	return &Zarinpal{
		cfg: cfg,
	}
}

func (z *Zarinpal) PaymentRequest(amount int, description string) (string, error) {
	// Construct the payment request URL
	paymentRequestUrl := z.cfg.Zarinpal.PaymentRequestUrl

	// Create the request body
	requestBody := map[string]interface{}{
		"merchant_id":  z.cfg.Zarinpal.MerchantId,
		"amount":       amount,
		"description":  description,
		"callback_url": z.cfg.Zarinpal.CallbackUrl,
	}

	// Make the HTTP POST request to Zarinpal's payment request endpoint
	response, err := makeHttpPostRequest(paymentRequestUrl, requestBody)
	if err != nil {
		return "", err
	}

	// Parse the response and return the authority
	responseDto, err := common.TypeConverter[PaymentRequestResponse](response)
	if err != nil {
		return "", err
	}

	return (z.cfg.Zarinpal.PaymentPageUrl + responseDto.Authority), nil

}

func makeHttpPostRequest(url string, body map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil

}
