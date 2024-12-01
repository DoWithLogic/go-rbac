package types_test

import (
	"bytes"
	"testing"

	"github.com/DoWithLogic/go-rbac/pkg/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestSecret_Masking(t *testing.T) {
	// Create a buffer to capture the log output
	var buf bytes.Buffer

	// Set up the logger to write to the buffer
	logger := zerolog.New(&buf).With().Timestamp().Logger()

	// Initialize the struct with test data
	loggingData := struct {
		Name     string       `json:"name"`
		Email    string       `json:"email"`
		Password types.Secret `json:"password"`
	}{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "your-secret-password",
	}

	// Log the struct
	logger.Info().Interface("data", loggingData).Msg("masking_password")

	// Assert that the log output matches the expected output
	assert.Contains(t, buf.String(), loggingData.Password.Masking())
}
