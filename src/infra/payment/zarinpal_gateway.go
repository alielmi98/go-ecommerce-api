package payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/alielmi98/go-ecommerce-api/common"
	"github.com/alielmi98/go-ecommerce-api/config"
)

// Zarinpal is a struct that implements the PaymentGateway interface
type ZarinpalResponse struct {
	Data   PaymentRequestResponse `json:"data,omitempty"`
	Errors json.RawMessage        `json:"errors,omitempty"`
}

// Payment Request Response Dto
type PaymentRequestResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Authority string `json:"authority"`
	RefId     int    `json:"ref_id"`
}

// payment Request Error Dto

type Zarinpal struct {
	cfg *config.Config
}

func NewZarinpalGateway(cfg *config.Config) *Zarinpal {
	return &Zarinpal{
		cfg: cfg,
	}
}

func (z *Zarinpal) PaymentRequest(amount float64, description string) (string, string, error) {
	// Construct the payment request URL
	var paymentRequestUrl string
	var paymentPageUrl string
	if z.cfg.Zarinpal.Sandbox {
		paymentRequestUrl = z.cfg.Zarinpal.SandboxPaymentRequestUrl
		paymentPageUrl = z.cfg.Zarinpal.SandboxPaymentPageUrl
	} else {
		paymentRequestUrl = z.cfg.Zarinpal.PaymentRequestUrl
		paymentPageUrl = z.cfg.Zarinpal.PaymentPageUrl
	}

	// Create the request body
	requestBody := map[string]interface{}{
		"merchant_id":  z.cfg.Zarinpal.MerchantId,
		"amount":       amount,
		"description":  description,
		"callback_url": "http://" + z.cfg.Server.Domain + ":" + string(z.cfg.Server.InternalPort) + z.cfg.Zarinpal.CallbackUrl,
	}

	// Make the HTTP POST request to Zarinpal's payment request endpoint
	zarinPalResponse, err := makeHttpPostRequest(paymentRequestUrl, requestBody)
	if err != nil {
		return "", "", err
	}

	// Unmarshal the response body into a map
	var response map[string]interface{}
	err = json.Unmarshal(zarinPalResponse, &response)
	if err != nil {
		return "", "", err
	}

	// Parse the response and return the authority
	responseDto, err := common.TypeConverter[ZarinpalResponse](response)
	if err != nil {
		return "", "", err
	}

	// Use the helper to extract error message
	if msg := extractErrorMessage(responseDto.Errors); msg != "" {
		return "", "", errors.New(msg)
	}
	if responseDto.Data.Code != 100 {
		return "", "", errors.New(responseDto.Data.Message)
	}
	return paymentPageUrl, responseDto.Data.Authority, nil

}
func (z *Zarinpal) PaymentVerify(authority string, amount float64) (int, error) {
	// Construct the payment verify URL
	var paymentVerifyUrl string
	if z.cfg.Zarinpal.Sandbox {
		paymentVerifyUrl = z.cfg.Zarinpal.SandboxPaymentVerifyUrl
	} else {
		paymentVerifyUrl = z.cfg.Zarinpal.PaymentVerifyUrl
	}

	// Create the request body
	requestBody := map[string]interface{}{
		"merchant_id": z.cfg.Zarinpal.MerchantId,
		"amount":      amount,
		"authority":   authority,
	}

	// Make the HTTP POST request to Zarinpal's payment verify endpoint
	zarinPalResponse, err := makeHttpPostRequest(paymentVerifyUrl, requestBody)
	if err != nil {
		return 0, err
	}

	// Unmarshal the response body into a map
	var response map[string]interface{}
	err = json.Unmarshal(zarinPalResponse, &response)
	if err != nil {
		return 0, err
	}

	responseDto, err := common.TypeConverter[ZarinpalResponse](response)
	if err != nil {
		return 0, err
	}

	// Use the helper to extract error message
	if msg := extractErrorMessage(responseDto.Errors); msg != "" {
		return 0, errors.New(msg)
	}
	if responseDto.Data.Code != 100 && responseDto.Data.Code != 101 {
		return 0, errors.New(responseDto.Data.Message)
	}
	return responseDto.Data.RefId, nil
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

// Helper function to extract "message" from RawMessage (object or array)
func extractErrorMessage(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	// Try object
	var obj struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(raw, &obj); err == nil && obj.Message != "" {
		return obj.Message
	}
	// Try array
	var arr []struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(raw, &arr); err == nil && len(arr) > 0 && arr[0].Message != "" {
		return arr[0].Message
	}
	return ""
}
