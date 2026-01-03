package scan

import (
	"fmt"
	"os"
	"sync"
	"yieldaa/cli/internal/info"
)

func ProcessEntities(files []info.EntityFile, maxWorkers int) ([]ScannedEntity, []error) {
	if len(files) == 0 {
		return []ScannedEntity{}, nil
	}

	if maxWorkers > len(files) || maxWorkers <= 0 {
		maxWorkers = len(files)
	}

	jobs := make(chan info.EntityFile, len(files))
	results := make(chan ScannedEntity, len(files))
	errors := make(chan error, len(files))

	var wg sync.WaitGroup
	var seenHashes sync.Map    // thread-safe для хешей контента
	var seenKeys sync.Map      // thread-safe для проверки ключей сущностей
	var keyConflicts []string  // для сбора конфликтов
	var conflictsMu sync.Mutex // мьютекс для keyConflicts

	// воркеры
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for file := range jobs {
				content, err := os.ReadFile(file.Path)
				if err != nil {
					errors <- fmt.Errorf("%s: read: %w", file.Path, err)
					continue
				}

				// xxHash64 вместо CRC32
				contentHash := calculateContentHash(content)

				// Atomic check and store
				if _, alreadyProcessed := seenHashes.LoadOrStore(contentHash, true); alreadyProcessed {
					continue
				}

				result := ProcessEntity(file, content)

				// check entity key
				if result.ParsedData != nil && result.FatalError == nil {
					key := EntityKey(result.ParsedData)
					if key != "" {
						if existingFile, exists := seenKeys.LoadOrStore(key, file.Path); exists {
							result.Errors = append(result.Errors,
								fmt.Sprintf("entity key conflict: '%s' already defined in '%s'",
									key, existingFile.(string)))

							conflictsMu.Lock()
							keyConflicts = append(keyConflicts,
								fmt.Sprintf("  %s:\n    • %s\n    • %s",
									key, existingFile.(string), file.Path))
							conflictsMu.Unlock()
						}
					}
				}

				if result.FatalError != nil {
					errors <- fmt.Errorf("%s: %w", file.Path, result.FatalError)
				} else {
					results <- result
				}
			}
		}(i)
	}

	// Feed jobs
	go func() {
		for _, file := range files {
			jobs <- file
		}
		close(jobs)
	}()

	// Wait completion
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	// Collect results
	var processed []ScannedEntity
	var fatalErrors []error

	for result := range results {
		processed = append(processed, result)
	}
	for err := range errors {
		fatalErrors = append(fatalErrors, err)
	}

	return processed, fatalErrors
}
