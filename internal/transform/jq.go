package transform

import (
	"errors"
	"log"

	"github.com/itchyny/gojq"
	"github.com/uk-devops/yaq/internal/pipeline"
)

func init() {
	RegisterTransformFunction("jq", ProcessWithJQ)
}

func ProcessWithJQ(inputData pipeline.StructuredData, jqExpression string) (pipeline.StructuredData, error) {
	query, err := gojq.Parse(jqExpression)
	if err != nil {
		log.Fatalln(err)
	}

	var iter gojq.Iter

	dataMap, ok := inputData.(pipeline.GenericMap)
	if !ok {
		dataArray, ok := inputData.(pipeline.GenericArray)
		if !ok {
			return nil, errors.New("not a map or array")
		}
		iter = query.Run([]interface{}(dataArray))
	} else {
		iter = query.Run(map[string]interface{}(dataMap))
	}

	var output pipeline.GenericArray

	for {
		i, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := i.(error); ok {
			log.Fatalln(err)
		}
		output = append(output, i)
	}

	if len(output) == 1 {
		switch o := output[0].(type) {
		case bool:
			return pipeline.GenericMap{"result": o}, nil
		case int:
			return pipeline.GenericMap{"result": o}, nil
		case string:
			return pipeline.GenericMap{"result": o}, nil
		case map[string]interface{}:
			return pipeline.GenericMap(o), nil
		case []interface{}:
			return pipeline.GenericArray(o), nil
		default:
			return nil, errors.New("cannot decode jq result")
		}
	} else {
		return output, nil
	}

}
