package common

import (
	"encoding/base64"
	"path/filepath"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/google/uuid"
)

func AddB64Resource(filePath string, data []byte, title, description string) (*oscalTypes.Resource, error) {
	
	encodedContent := base64.StdEncoding.EncodeToString(data)

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