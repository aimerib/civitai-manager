package services

import (
	"civitai-manager/models"
	"regexp"

	"gorm.io/gorm"
)

type ModelProcessor struct {
	DB *gorm.DB
}

func NewModelProcessor(db *gorm.DB) *ModelProcessor {
	return &ModelProcessor{DB: db}
}

func (mp *ModelProcessor) ProcessModel(model *models.Model) error {
	return mp.DB.Transaction(func(tx *gorm.DB) error {
		if model.Description != "" {
			re := regexp.MustCompile(`\s*id="[^"]*"`)
			model.Description = re.ReplaceAllString(model.Description, "")
		}

		if err := tx.Create(model).Error; err != nil {
			return err
		}

		if err := mp.ProcessModelVersions(tx, model); err != nil {
			return err
		}

		if err := mp.ProcessTags(tx, model); err != nil {
			return err
		}

		if err := mp.ProcessStats(tx, model); err != nil {
			return err
		}

		if err := mp.ProcessCreator(tx, model); err != nil {
			return err
		}

		return nil
	})
}

func (mp *ModelProcessor) ProcessModelVersions(tx *gorm.DB, model *models.Model) error {
	for i := range model.ModelVersions {
		version := &model.ModelVersions[i]
		version.ModelID = model.ID

		if err := tx.Create(version).Error; err != nil {
			return err
		}

		if err := mp.ProcessFiles(tx, version); err != nil {
			return err
		}

		if err := mp.processImages(tx, version); err != nil {
			return err
		}

		if err := mp.processModelVersionStats(tx, version); err != nil {
			return err
		}
	}
	return nil
}

func (mp *ModelProcessor) ProcessTags(tx *gorm.DB, model *models.Model) error {
	for _, tag := range model.Tags {
		var existingTag models.Tag
		err := tx.Where("name = ?", tag.Name).FirstOrCreate(&existingTag).Error
		if err != nil {
			return err
		}

		if err := tx.Model(model).Association("Tags").Append(&existingTag); err != nil {
			return err
		}
	}
	return nil
}

func (mp *ModelProcessor) ProcessStats(tx *gorm.DB, model *models.Model) error {
	model.Stats.ModelID = model.ID
	return tx.Create(&model.Stats).Error
}

func (mp *ModelProcessor) ProcessCreator(tx *gorm.DB, model *models.Model) error {
	model.Creator.ModelID = model.ID
	return tx.Create(&model.Creator).Error
}

func (mp *ModelProcessor) ProcessFiles(tx *gorm.DB, modelVersion *models.ModelVersion) error {
	for i := range modelVersion.Files {
		file := &modelVersion.Files[i]
		file.ModelVersionID = modelVersion.ID
		if err := tx.Create(file).Error; err != nil {
			return err
		}
	}
	return nil
}

func (mp *ModelProcessor) processImages(tx *gorm.DB, modelVersion *models.ModelVersion) error {
	for i := range modelVersion.Images {
		image := &modelVersion.Images[i]
		image.ModelVersionID = modelVersion.ID
		if err := tx.Create(image).Error; err != nil {
			return err
		}
	}
	return nil
}

func (mp *ModelProcessor) processModelVersionStats(tx *gorm.DB, modelVersion *models.ModelVersion) error {
	modelVersion.ModelVersionStat.ModelVersionID = modelVersion.ID
	return tx.Create(&modelVersion.ModelVersionStat).Error
}
