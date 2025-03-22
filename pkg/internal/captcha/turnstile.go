package captcha

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type TurnstileAdapter struct{}

type turnstileResponse struct {
	Success    bool     `json:"success"`
	ErrorCodes []string `json:"error-codes"`
}

func (a *TurnstileAdapter) Validate(token, ip string) bool {
	url := "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	data := map[string]string{
		"secret":   viper.GetString("captcha.api_secret"),
		"response": token,
		"remoteip": ip,
	}

	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("Error sending request to Turnstile...")
		return false
	}
	defer resp.Body.Close()

	var result turnstileResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Error().Err(err).Msg("Error decoding response from Turnstile...")
		return false
	}

	if !result.Success {
		log.Warn().Strs("errors", result.ErrorCodes).Msg("An captcha validation request failed...")
	}

	return result.Success
}
