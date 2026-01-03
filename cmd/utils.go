package cmd

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	// Для коротких интервалов показываем миллисекунды
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}

	// Для средних - секунды с миллисекундами
	if d < 10*time.Second {
		return fmt.Sprintf("%.3fs", d.Seconds())
	}

	// Для длинных - секунды
	return fmt.Sprintf("%.1fs", d.Seconds())
}

// Или готовый вариант с подсказками:
func formatDurationHuman(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return fmt.Sprintf("%dµs", d.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%dms", d.Milliseconds())
	case d < time.Minute:
		secs := d.Seconds()
		if secs < 10 {
			return fmt.Sprintf("%.3fs", secs) // 1.234s
		}
		return fmt.Sprintf("%.1fs", secs) // 12.3s
	case d < time.Hour:
		return fmt.Sprintf("%.1fm", d.Minutes())
	default:
		return fmt.Sprintf("%.1fh", d.Hours())
	}
}
