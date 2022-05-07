package pipeline

import "fmt"

// GenericMap is a generic map with keys as strings
type GenericMap map[string]interface{}

// GenericArray is a generic array of interfaces
type GenericArray []interface{}

// StructuredData is a container of either a GenericMap or GenericArray
type StructuredData struct {
	Data interface{}
}

func (d *StructuredData) onlyMaps(newData interface{}) (GenericMap, GenericMap, bool) {
	m1, ok1 := d.Data.(GenericMap)
	m2, ok2 := newData.(GenericMap)
	if ok1 && ok2 {
		return m1, m2, true
	}
	return nil, nil, false
}

func (d *StructuredData) onlyArrays(newData interface{}) (GenericArray, GenericArray, bool) {
	a1, ok1 := d.Data.(GenericArray)
	a2, ok2 := newData.(GenericArray)
	if ok1 && ok2 {
		return a1, a2, true
	}
	return nil, nil, false
}

func (d *StructuredData) mapAndArray(newData interface{}) (GenericMap, GenericArray, bool) {
	m1, ok1 := d.Data.(GenericMap)
	a2, ok2 := newData.(GenericArray)
	if ok1 && ok2 {
		return m1, a2, true
	}
	return nil, nil, false
}

func (d *StructuredData) arrayAndMap(newData interface{}) (GenericArray, GenericMap, bool) {
	a1, ok1 := d.Data.(GenericArray)
	m2, ok2 := newData.(GenericMap)
	if ok1 && ok2 {
		return a1, m2, true
	}
	return nil, nil, false
}

func (d *StructuredData) dataIsEmpty() bool {
	return d.Data == nil
}

func (d *StructuredData) Append(newData interface{}) {
	if d.dataIsEmpty() {
		d.Data = newData
	} else if m1, m2, ok := d.onlyMaps(newData); ok {
		d.Data = mergeMaps(m1, m2)
	} else if a1, a2, ok := d.onlyArrays(newData); ok {
		d.Data = append(a1, a2...)
	} else if m1, a2, ok := d.mapAndArray(newData); ok {
		d.Data = append(GenericArray{m1}, a2...)
	} else if a1, m2, ok := d.arrayAndMap(newData); ok {
		d.Data = append(a1, GenericArray{m2}...)
	} else {
		panic(fmt.Sprintf("Can't append %v to %v", newData, d.Data))
	}
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
