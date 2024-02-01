package integrationruntime

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SsisObjectMetadata interface {
}

// RawSsisObjectMetadataImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSsisObjectMetadataImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalSsisObjectMetadataImplementation(input []byte) (SsisObjectMetadata, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SsisObjectMetadata into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Environment") {
		var out SsisEnvironment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SsisEnvironment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Folder") {
		var out SsisFolder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SsisFolder: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Package") {
		var out SsisPackage
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SsisPackage: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Project") {
		var out SsisProject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SsisProject: %+v", err)
		}
		return out, nil
	}

	out := RawSsisObjectMetadataImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
