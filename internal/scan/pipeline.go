package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	PACKAGE_ENTITIES_DIR_NAME = "entities"
)

func LoadPackage(wd string) (*Package, error) {
	var packageData Package // данные пакета

	// загрузка конфигурации
	config, err := ReadConfig(wd)
	if err != nil {
		return nil, fmt.Errorf("error read package configuration: %v\n", err)
	}
	packageData.Config = *config

	// определение директории сущностей пакета
	entitiesDir := filepath.Join(wd, PACKAGE_ENTITIES_DIR_NAME)
	if _, err := os.Stat(entitiesDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("package does not have '%s' directory\n", PACKAGE_ENTITIES_DIR_NAME)
	}

	// сканирование содержимого пакета
	entitiesFiles, err := ScanEntities(entitiesDir)
	if err != nil {
		return nil, fmt.Errorf("error read package configuration: %v\n", err)
	}
	packageData.Files = entitiesFiles

	// результат
	totalSize := int64(0)
	for _, f := range packageData.Files {
		totalSize += f.Size
	}
	packageData.Sum.TotalSize = totalSize
	packageData.Sum.EntitiesCount = len(packageData.Files)
	packageData.Sum.ControlHash = CalculateStructureHash(packageData.Files)

	return &packageData, nil
}
