package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type File struct {
	ID                int             `json:"-" db:"id"`
	CivitaiID         *int            `json:"id" db:"civitai_id"`
	ModelVersionsID   int             `json:"-" db:"model_versions_id"`
	SizeKB            float64         `json:"sizeKB" db:"size_kb"`
	Name              *string         `json:"name" db:"name"`
	Type              *string         `json:"type" db:"type"`
	PickleScanResult  *string         `json:"pickleScanResult" db:"pickle_scan_result"`
	PickleScanMessage *string         `json:"pickleScanMessage" db:"pickle_scan_message"`
	VirusScanResult   *string         `json:"virusScanResult" db:"virus_scan_result"`
	VirusScanMessage  *string         `json:"virusScanMessage" db:"virus_scan_message"`
	ScannedAt         time.Time       `json:"scannedAt" db:"scanned_at"`
	Metadata          json.RawMessage `json:"metadata" db:"metadata"`
	Hashes            json.RawMessage `json:"hashes" db:"hashes"`
	DownloadURL       string          `json:"downloadUrl" db:"download_url"`
	IsPrimary         *bool           `json:"primary" db:"is_primary"`
	CreatedAt         time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" db:"updated_at"`
	ModelVersions     ModelVersions   `json:"-" belongs_to:"model_versions"`
}

type Files []File

type FileMetadata struct {
	FP     string `json:"fp"`
	Size   string `json:"size"`
	Format string `json:"format"`
}

type Hashes struct {
	AutoV2 string `json:"AutoV2"`
	SHA256 string `json:"SHA256"`
	CRC32  string `json:"CRC32"`
	BLAKE3 string `json:"BLAKE3"`
}

func (f *File) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
