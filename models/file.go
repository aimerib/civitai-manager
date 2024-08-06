package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID                uint            `gorm:"primaryKey" json:"-"`
	CivitaiID         *int            `json:"id"`
	ModelVersionID    uint            `gorm:"index" json:"-"`
	SizeKB            float64         `json:"sizeKB"`
	Name              *string         `json:"name"`
	Type              *string         `json:"type"`
	PickleScanResult  *string         `json:"pickleScanResult"`
	PickleScanMessage *string         `json:"pickleScanMessage"`
	VirusScanResult   *string         `json:"virusScanResult"`
	VirusScanMessage  *string         `json:"virusScanMessage"`
	ScannedAt         time.Time       `json:"scannedAt"`
	Metadata          json.RawMessage `gorm:"type:json" json:"metadata"`
	Hashes            json.RawMessage `gorm:"type:json" json:"hashes"`
	DownloadURL       string          `json:"downloadUrl"`
	IsPrimary         *bool           `json:"primary"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

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

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return
}

func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	f.UpdatedAt = time.Now()
	return
}

func (f *File) GetMetadata() (FileMetadata, error) {
	var metadata FileMetadata
	err := json.Unmarshal(f.Metadata, &metadata)
	return metadata, err
}

func (f *File) GetHashes() (Hashes, error) {
	var hashes Hashes
	err := json.Unmarshal(f.Hashes, &hashes)
	return hashes, err
}
