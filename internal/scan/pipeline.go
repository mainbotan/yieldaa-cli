package scan

import (
	"fmt"
	"yieldaa/cli/internal/info"
)

const (
	MAX_WORKERS = 100
)

func ScanEntities(wd string) ([]ScannedEntity, error) {
	packageData, err := info.LoadPackage(wd)
	if err != nil {
		return nil, fmt.Errorf("error load package: %w", err)
	}

	result, processErrors := ProcessEntities(packageData.Files, MAX_WORKERS)
	if len(processErrors) > 0 {
		// Конвертируем []error в одну ошибку
		return nil, fmt.Errorf("processing errors: %v", processErrors)
	}

	return result, nil
}
