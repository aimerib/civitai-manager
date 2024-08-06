package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var sensitiveParams = []string{"password", "token", "api_key"}

func ParamLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log URL parameters
		if len(c.Params) > 0 {
			params := make(map[string]string)
			for _, p := range c.Params {
				params[p.Key] = filterSensitiveData(p.Key, p.Value)
			}
			logger.Info("URL Parameters", zap.Any("params", params))
		}

		// Log query parameters
		if len(c.Request.URL.Query()) > 0 {
			query := filterQueryParams(c.Request.URL.Query())
			logger.Info("Query Parameters", zap.Any("query", query))
		}

		// Log form data
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			if err := c.Request.ParseForm(); err == nil {
				formData := filterQueryParams(c.Request.PostForm)
				logger.Info("Form Data", zap.Any("form", formData))
			}
		}

		// Log JSON body
		if c.Request.Header.Get("Content-Type") == "application/json" {
			body, err := io.ReadAll(c.Request.Body)
			if err == nil {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
				var jsonMap map[string]interface{}
				if json.Unmarshal(body, &jsonMap) == nil {
					filteredJSON := filterJSONParams(jsonMap)
					logger.Info("JSON Body", zap.Any("json", filteredJSON))
				}
			}
		}

		c.Next()
	}
}

func filterSensitiveData(key, value string) string {
	for _, param := range sensitiveParams {
		if strings.ToLower(key) == param {
			return "[FILTERED]"
		}
	}
	return value
}

func filterQueryParams(params url.Values) url.Values {
	filtered := make(url.Values)
	for key, values := range params {
		filteredValues := make([]string, len(values))
		for i, value := range values {
			filteredValues[i] = filterSensitiveData(key, value)
		}
		filtered[key] = filteredValues
	}
	return filtered
}

func filterJSONParams(jsonMap map[string]interface{}) map[string]interface{} {
	filtered := make(map[string]interface{})
	for key, value := range jsonMap {
		switch v := value.(type) {
		case string:
			filtered[key] = filterSensitiveData(key, v)
		case map[string]interface{}:
			filtered[key] = filterJSONParams(v)
		default:
			filtered[key] = v
		}
	}
	return filtered
}
