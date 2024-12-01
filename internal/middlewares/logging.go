package middlewares

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

// Implement Write method
func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	// Write to both the buffer and the original writer
	n, err := crw.buf.Write(b)
	if err != nil {
		return n, err
	}
	return crw.ResponseWriter.Write(b)
}

// Implement Hijack method to support WebSocket connections
func (crw *CustomResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	// Ensure that the original ResponseWriter supports Hijacker
	hijacker, ok := crw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("ResponseWriter does not implement http.Hijacker")
	}
	return hijacker.Hijack()
}

var ignoredPatterns = []*regexp.Regexp{}

func init() {
	// Define the dynamic parts of the paths
	dynamicPaths := []string{"swagger", "ping", "metrics"}

	for _, path := range dynamicPaths {
		// Create a pattern that allows for any prefix before the dynamic path
		pattern := fmt.Sprintf(`.*/%s`, path)
		ignoredPatterns = append(ignoredPatterns, regexp.MustCompile(pattern))
	}
	// Additional static pattern for /swagger/*
	ignoredPatterns = append(ignoredPatterns, regexp.MustCompile(`^/swagger/`))
}

func (m *Middleware) isPathIgnored(path string) bool {
	for _, pattern := range ignoredPatterns {
		if pattern.MatchString(path) {
			return true
		}
	}
	return false
}

var maskedKeys = []string{"password", "token"}

func (m *Middleware) isMaskedKey(key string) bool {
	for _, maskedKey := range maskedKeys {
		if strings.Contains(key, maskedKey) {
			return true
		}
	}

	return false
}

// Helper function to mask sensitive fields
func (m *Middleware) maskSensitiveFields(data map[string]interface{}) map[string]interface{} {
	for key, value := range data {
		// Check if the key matches any sensitive key patterns in maskedKeys
		if m.isMaskedKey(key) {
			data[key] = "*****"
		} else if nested, ok := value.(map[string]interface{}); ok {
			// Recursively mask nested fields
			data[key] = m.maskSensitiveFields(nested)
		}
	}
	return data
}

func (m *Middleware) LoggingMiddleware(logger *zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			startTime := time.Now()

			ignoreLogging := m.isPathIgnored(req.URL.Path)

			// Get or generate X-Request-ID
			requestID := req.Header.Get(constants.XRequestIDKey.String())
			if requestID == "" {
				requestID = uuid.New().String()
				res.Header().Set(constants.XRequestIDKey.String(), requestID)
			}

			var requestBodyJSON map[string]interface{}
			if req.Body != nil {
				// Read and restore the request body
				reqBody, err := io.ReadAll(req.Body)
				if err == nil {
					req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

					// Parse the request body as JSON
					if err := json.Unmarshal(reqBody, &requestBodyJSON); err == nil {
						// Mask sensitive fields
						requestBodyJSON = m.maskSensitiveFields(requestBodyJSON)
					}
				}
			}

			c.SetRequest(c.Request().WithContext(
				context.WithValue(
					c.Request().Context(),
					constants.XRequestIDKey,
					requestID,
				)))

			// Create a custom response writer to capture the response body
			buf := new(bytes.Buffer)
			crw := &CustomResponseWriter{ResponseWriter: c.Response().Writer, buf: buf}
			c.Response().Writer = crw

			// Continue to the next middleware/handler
			err := next(c)

			stop := time.Now()
			latency := stop.Sub(startTime)

			if !ignoreLogging {
				cl := req.Header.Get(echo.HeaderContentLength)
				if cl == "" {
					cl = "0"
				}

				logEvent := logger.Info().
					Ctx(c.Request().Context()).
					Str("remote_ip", req.RemoteAddr).
					Str("host", req.Host).
					Str("method", req.Method).
					Str("uri", req.RequestURI).
					Str("user_agent", req.UserAgent()).
					Int("status", res.Status).
					Str("referer", req.Referer()).
					Dur("latency", latency).
					Str("latency_human", latency.String()).
					Str("bytes_in", cl).
					Str("bytes_out", strconv.FormatInt(res.Size, 10))

				// Check if the response is JSON
				if strings.Contains(res.Header().Get(echo.HeaderContentType), echo.MIMEApplicationJSON) {
					var responseBodyJSON map[string]interface{}
					if err := json.Unmarshal(buf.Bytes(), &responseBodyJSON); err == nil {
						responseBodyJSON = m.maskSensitiveFields(responseBodyJSON)
					}

					logEvent.Interface("request_body", requestBodyJSON)
					logEvent.Interface("response_body", responseBodyJSON)
				}

				logEvent.Msg("[Received HTTP request]")
			}

			return err
		}
	}
}
