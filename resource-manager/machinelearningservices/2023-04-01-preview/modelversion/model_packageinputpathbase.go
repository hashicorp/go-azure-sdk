package modelversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageInputPathBase interface {
}

func unmarshalPackageInputPathBaseImplementation(input []byte) (PackageInputPathBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PackageInputPathBase into map[string]interface: %+v", err)
	}

	value, ok := temp["inputPathType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "PathId") {
		var out PackageInputPathId
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathId: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Url") {
		var out PackageInputPathUrl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathUrl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PathVersion") {
		var out PackageInputPathVersion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PackageInputPathVersion: %+v", err)
		}
		return out, nil
	}

	type RawPackageInputPathBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawPackageInputPathBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
