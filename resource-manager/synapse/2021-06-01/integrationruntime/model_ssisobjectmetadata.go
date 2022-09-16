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

	type RawSsisObjectMetadataImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawSsisObjectMetadataImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
