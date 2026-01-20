package int

func ToInt(value any, fallback int) int {
	switch cast := value.(type) {
	case int:
		return cast
	case int32:
		return int(cast)
	case int64:
		return int(cast)
	case float32:
		return int(cast)
	case float64:
		return int(cast)
	default:
		return fallback
	}
}
