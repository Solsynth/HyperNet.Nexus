package captcha

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type TemplateData struct {
	ApiKey string `json:"api_key"`
}

func GetTemplateData() TemplateData {
	return TemplateData{
		ApiKey: viper.GetString("captcha.api_key"),
	}
}

type CaptchaAdapter interface {
	Validate(token, ip string) bool
}

var adapters = map[string]CaptchaAdapter{
	"turnstile": &TurnstileAdapter{},
}

func Validate(token, ip string) bool {
	provider := viper.GetString("captcha.provider")
	if adapter, ok := adapters[provider]; ok {
		return adapter.Validate(token, ip)
	}
	log.Error().Msg("Unable to handle captcha validate request due to unsupported provider.")
	return false
}
