package dataversionregistry

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataImportSource interface {
}

// RawDataImportSourceImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataImportSourceImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalDataImportSourceImplementation(input []byte) (DataImportSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataImportSource into map[string]interface: %+v", err)
	}

	value, ok := temp["sourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "database") {
		var out DatabaseSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatabaseSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "file_system") {
		var out FileSystemSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileSystemSource: %+v", err)
		}
		return out, nil
	}

	out := RawDataImportSourceImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
