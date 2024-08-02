package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
)

type File struct {
	ID                int          `db:"id"`
	CivitaiID         int          `json:"id" db:"civitai_id"`
	ModelVersionID    int          `db:"model_version_id"`
	SizeKB            float64      `json:"sizeKB" db:"size_kb"`
	Name              string       `json:"name" db:"name"`
	Type              string       `json:"type" db:"type"`
	PickleScanResult  *string      `json:"pickleScanResult" db:"pickle_scan_result"`
	PickleScanMessage *string      `json:"pickleScanMessage" db:"pickle_scan_message"`
	VirusScanResult   *string      `json:"virusScanResult" db:"virus_scan_result"`
	VirusScanMessage  *string      `json:"virusScanMessage" db:"virus_scan_message"`
	ScannedAt         *time.Time   `json:"scannedAt" db:"scanned_at"`
	Metadata          FileMetadata `json:"metadata" db:"metadata"`
	Hashes            Hashes       `json:"hashes" db:"hashes"`
	DownloadURL       *string      `json:"downloaUrl" db:"download_url"`
	Primary           bool         `json:"primary" db:"primary"`
	CreatedAt         time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at" db:"updated_at"`
}

type Files []File

type FileMetadata struct {
	FP     *string `json:"fp"`
	Size   *string `json:"size"`
	Format *string `json:"format"`
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
