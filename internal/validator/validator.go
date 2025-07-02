package validator

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) IsValid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	_, exists := v.FieldErrors[key]
	if !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) ValidateNotBlank(val, key, message string) {
	if strings.TrimSpace(val) == "" {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) ValidateLength(val string, max int, key, message string) {
	if utf8.RuneCountInString(val) > max {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) ValidateUrl(urlStr string, fieldName string) string {
	if strings.TrimSpace(urlStr) == "" {
		return ""
	}

	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		urlStr = "https://" + urlStr
	}

	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		v.AddFieldError(fieldName, "Invalid URL")
		return urlStr
	}

	return parsedUrl.String()
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}
