package addons

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Addon interface {
}

func unmarshalAddonImplementation(input []byte) (Addon, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Addon into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ArcForKubernetes") {
		var out ArcAddon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ArcAddon: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "IotEdge") {
		var out IoTAddon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IoTAddon: %+v", err)
		}
		return out, nil
	}

	type RawAddonImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAddonImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
