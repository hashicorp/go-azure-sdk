package incidententities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Entity interface {
}

func unmarshalEntityImplementation(input []byte) (Entity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Entity into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bookmark") {
		var out HuntingBookmark
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HuntingBookmark: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SecurityAlert") {
		var out SecurityAlert
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAlert: %+v", err)
		}
		return out, nil
	}

	type RawEntityImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawEntityImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
