package pipeline

import (
	"fmt"
)

// GenericMap is a generic map with keys as strings
type GenericMap map[string]interface{}

// GenericArray is a generic array of interfaces
type GenericArray []interface{}

type StructuredData interface {
	Append(StructuredData) StructuredData
	Map() GenericMap
}

func (d GenericMap) Append(newData StructuredData) StructuredData {
	if m2, ok := newData.(GenericMap); ok {
		m := mergeMaps(d, m2)
		return m
	} else if a2, ok := newData.(GenericArray); ok {
		a := append(GenericArray{d}, a2...)
		return a
	} else {
		panic(fmt.Sprintf("Can't append %v to %v", newData, d))
	}
}

func (d GenericMap) Map() GenericMap {
	return d
}

func (d GenericArray) Append(newData StructuredData) StructuredData {
	if a2, ok := newData.(GenericArray); ok {
		a := append(d, a2...)
		return a
	} else if m2, ok := newData.(GenericMap); ok {
		a := append(d, GenericArray{m2}...)
		return a
	} else {
		panic(fmt.Sprintf("Can't append %v to %v", newData, d))
	}
}

func (d GenericArray) Map() GenericMap {
	return nil
}

func mergeMaps(m1, m2 GenericMap) GenericMap {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

type UsageError struct {
	Message string
}

func (e *UsageError) Error() string {
	return e.Message
}
