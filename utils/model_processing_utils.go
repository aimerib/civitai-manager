package modelutils

import (
	"civitai/models"
	"regexp"

	"github.com/gobuffalo/pop/v6"
)

func ProcessModel(tx *pop.Connection, model *models.Model) error {
	if model.Description != nil {
		re := regexp.MustCompile(`\s*id="[^"]*"`)
		*model.Description = re.ReplaceAllString(*model.Description, "")
	}

	err := tx.Create(model)
	if err != nil {
		return err
	}

	// Process associated data
	err = ProcessModelVersions(tx, model)
	if err != nil {
		return err
	}

	err = ProcessTags(tx, model)
	if err != nil {
		return err
	}

	err = ProcessStats(tx, model)
	if err != nil {
		return err
	}

	err = ProcessCreator(tx, model)
	if err != nil {
		return err
	}

	return nil
}

func ProcessModelVersions(tx *pop.Connection, model *models.Model) error {
	for _, version := range model.ModelVersions {
		version.ModelID = model.ID
		err := tx.Create(&version)
		if err != nil {
			return err
		}

		err = ProcessFiles(tx, &version)
		if err != nil {
			return err
		}

		err = processImages(tx, &version)
		if err != nil {
			return err
		}

		err = processModelVersionStats(tx, &version)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessTags(tx *pop.Connection, model *models.Model) error {
	for _, tag := range model.Tags {
		err := tag.FindOrCreate(tx)
		if err != nil {
			return err
		}

		modelTag := models.ModelTag{
			ModelID: model.ID,
			TagID:   tag.ID,
		}
		err = tx.Create(&modelTag)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessStats(tx *pop.Connection, model *models.Model) error {
	model.Stats.ModelID = model.ID
	// Process model stats
	return tx.Create(&model.Stats)
}

func ProcessCreator(tx *pop.Connection, model *models.Model) error {
	model.Creator.ModelID = model.ID
	return tx.Create(&model.Creator)
}

func ProcessFiles(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	for _, file := range modelVersion.Files {
		file.ModelVersionsID = modelVersion.ID
		err := tx.Create(&file)
		if err != nil {
			return err
		}
	}
	return nil
}

func processImages(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	for _, image := range modelVersion.Images {
		image.ModelVersionsID = modelVersion.ID
		err := tx.Create(&image)
		if err != nil {
			return err
		}
	}
	return nil
}

func processModelVersionStats(tx *pop.Connection, modelVersion *models.ModelVersions) error {
	modelVersion.Stats.ModelVersionsID = modelVersion.ID
	return tx.Create(&modelVersion.Stats)
}
