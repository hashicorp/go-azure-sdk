package pipelines

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompressionReadSettings interface {
}

// RawCompressionReadSettingsImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCompressionReadSettingsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalCompressionReadSettingsImplementation(input []byte) (CompressionReadSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CompressionReadSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	out := RawCompressionReadSettingsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
