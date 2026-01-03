package info

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cespare/xxhash/v2"
)

type Package struct {
	Config Config       `json:"config"`
	Files  []EntityFile `json:"files"`
	Sum    PackageSum   `json:"sum"`
}

type PackageSum struct {
	TotalSize     int64  `json:"total_size"`
	EntitiesCount int    `json:"entities_count"`
	ControlHash   uint32 `json:"control_hash"`
}

type EntityFile struct {
	Path        string    `json:"path"`
	Size        int64     `json:"size"`
	ModTime     time.Time `json:"mod_time"`
	ContentHash string    `json:"content_hash,omitempty"`
}

// обход сущностей без чтения содержимого
func ScanEntities(dir string) ([]EntityFile, error) {
	var files []EntityFile

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(strings.ToLower(info.Name()), ".yml") &&
			!strings.HasSuffix(strings.ToLower(info.Name()), ".yaml") {
			return nil
		}

		files = append(files, EntityFile{
			Path:    path,
			Size:    info.Size(),
			ModTime: info.ModTime(),
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return files, nil
}

// контрольная сумма пакета
func CalculateStructureHash(files []EntityFile) uint32 {
	hash := xxhash.New()
	for _, f := range files {
		relPath := strings.TrimPrefix(f.Path, filepath.Dir(f.Path)+string(os.PathSeparator))
		hash.Write([]byte(relPath))
		binary.Write(hash, binary.LittleEndian, f.Size)
		binary.Write(hash, binary.LittleEndian, f.ModTime.Unix())
	}

	result := binary.LittleEndian.Uint32(hash.Sum(nil)[:4])
	return result
}
