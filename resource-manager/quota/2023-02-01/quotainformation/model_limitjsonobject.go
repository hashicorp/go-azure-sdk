package quotainformation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LimitJsonObject interface {
}

// RawLimitJsonObjectImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLimitJsonObjectImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalLimitJsonObjectImplementation(input []byte) (LimitJsonObject, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LimitJsonObject into map[string]interface: %+v", err)
	}

	value, ok := temp["limitObjectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "LimitValue") {
		var out LimitObject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LimitObject: %+v", err)
		}
		return out, nil
	}

	out := RawLimitJsonObjectImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
