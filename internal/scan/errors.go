package scan

import (
	"fmt"
)

// Типы ошибок по критичности
type ErrorSeverity int

const (
	SeverityFatal   ErrorSeverity = iota // Нельзя продолжать
	SeverityError                        // Серьёзная ошибка, но можно продолжать
	SeverityWarning                      // Предупреждение
	SeverityInfo                         // Информация
)

// Структурированная ошибка
type ScanError struct {
	FilePath   string
	Message    string
	Severity   ErrorSeverity
	EntityCode string // если применимо
	FieldCode  string // если применимо
}

func (e ScanError) Error() string {
	return fmt.Sprintf("%s: %s", e.FilePath, e.Message)
}

// Результат сканирования с классифицированными ошибками
type ScanResult struct {
	Entities     []ScannedEntity
	FatalErrors  []ScanError   // Критические - остановка сканирования
	Errors       []ScanError   // Ошибки валидации
	Warnings     []ScanError   // Предупреждения
	KeyConflicts []KeyConflict // Конфликты ключей
}

type KeyConflict struct {
	Key       string
	FilePaths []string // Все файлы, где встречается этот ключ
}
