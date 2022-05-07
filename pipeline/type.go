package pipeline

// GenericMap is a generic map with keys as strings
type GenericMap map[string]interface{}
type UsageError struct {
	Message string
}

func (e *UsageError) Error() string {
	return e.Message
}
