package common

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/google/uuid"
)

// AddB64Resource creates a back-matter resource with base64 encoded content
func AddB64Resource(filePath, title, description string) (*oscalTypes.Resource, error) {
	// Read the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file for base64 encoding: %w", err)
	}

	// Generate base64 encoded content
	encodedContent := base64.StdEncoding.EncodeToString(fileData)

	// Create resource
	resource := &oscalTypes.Resource{
		UUID:  uuid.New().String(),
		Title: title,
		Description: description,
		Base64: &oscalTypes.Base64{
			Filename:  filepath.Base(filePath),
			MediaType: "application/json",
			Value:     encodedContent,
		},
	}

	return resource, nil
}