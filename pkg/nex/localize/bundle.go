package localize

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	htmpl "html/template"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Bundle struct {
	Bundle *i18n.Bundle

	LocalesPath   string
	TemplatesPath string
}

const FallbackLanguage = "en-US"

var L *Bundle

func LoadLocalization(localesPath string, templatesPath ...string) error {
	L = &Bundle{
		LocalesPath: localesPath,
	}
	if len(templatesPath) > 0 {
		L.TemplatesPath = templatesPath[0]
	}

	L.Bundle = i18n.NewBundle(language.AmericanEnglish)
	L.Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var count int

	basePath := localesPath
	if entries, err := os.ReadDir(basePath); err != nil {
		return fmt.Errorf("unable to read locales directory: %v", err)
	} else {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if _, err := L.Bundle.LoadMessageFile(filepath.Join(basePath, entry.Name())); err != nil {
				return fmt.Errorf("unable to load localization file %s: %v", entry.Name(), err)
			} else {
				count++
			}
		}
	}

	log.Info().Int("locales", count).Msg("Loaded localization files...")

	return nil
}

func (v *Bundle) GetLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(v.Bundle, lang)
}

func (v *Bundle) GetLocalizedString(name string, lang string) string {
	localizer := v.GetLocalizer(lang)
	msg, err := localizer.LocalizeMessage(&i18n.Message{
		ID: name,
	})
	if err != nil {
		log.Warn().Err(err).Str("lang", lang).Str("name", name).Msg("Failed to localize string...")
		return name
	}
	return msg
}

func (v *Bundle) GetLocalizedTemplatePath(name string, lang string) string {
	basePath := v.TemplatesPath
	filePath := filepath.Join(basePath, lang, name)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// Fallback to English
		filePath = filepath.Join(basePath, FallbackLanguage, name)
		return filePath
	}

	return filePath
}

func (v *Bundle) GetLocalizedTemplate(name string, lang string) *template.Template {
	path := v.GetLocalizedTemplatePath(name, lang)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Warn().Err(err).Str("lang", lang).Str("name", name).Msg("Failed to load localized template...")
		return nil
	}

	return tmpl
}

func (v *Bundle) GetLocalizedTemplateHTML(name string, lang string) *htmpl.Template {
	path := v.GetLocalizedTemplatePath(name, lang)
	tmpl, err := htmpl.ParseFiles(path)
	if err != nil {
		log.Warn().Err(err).Str("lang", lang).Str("name", name).Msg("Failed to load localized template...")
		return nil
	}

	return tmpl
}

func (v *Bundle) RenderLocalizedTemplateHTML(name string, lang string, data any) string {
	tmpl := v.GetLocalizedTemplate(name, lang)
	if tmpl == nil {
		return ""
	}
	buf := new(strings.Builder)
	err := tmpl.Execute(buf, data)
	if err != nil {
		log.Warn().Err(err).Str("lang", lang).Str("name", name).Msg("Failed to render localized template...")
		return ""
	}
	return buf.String()
}

func (v *Bundle) RenderLocalizedTemplate(name string, lang string, data any) string {
	tmpl := v.GetLocalizedTemplate(name, lang)
	if tmpl == nil {
		return ""
	}
	buf := new(strings.Builder)
	err := tmpl.Execute(buf, data)
	if err != nil {
		log.Warn().Err(err).Str("lang", lang).Str("name", name).Msg("Failed to render localized template...")
		return ""
	}
	return buf.String()
}
