package quotarequeststatus

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LimitJsonObject interface {
	LimitJsonObject() BaseLimitJsonObjectImpl
}

var _ LimitJsonObject = BaseLimitJsonObjectImpl{}

type BaseLimitJsonObjectImpl struct {
	LimitObjectType LimitType `json:"limitObjectType"`
}

func (s BaseLimitJsonObjectImpl) LimitJsonObject() BaseLimitJsonObjectImpl {
	return s
}

var _ LimitJsonObject = RawLimitJsonObjectImpl{}

// RawLimitJsonObjectImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLimitJsonObjectImpl struct {
	limitJsonObject BaseLimitJsonObjectImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawLimitJsonObjectImpl) LimitJsonObject() BaseLimitJsonObjectImpl {
	return s.limitJsonObject
}

func UnmarshalLimitJsonObjectImplementation(input []byte) (LimitJsonObject, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LimitJsonObject into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["limitObjectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "LimitValue") {
		var out LimitObject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LimitObject: %+v", err)
		}
		return out, nil
	}

	var parent BaseLimitJsonObjectImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLimitJsonObjectImpl: %+v", err)
	}

	return RawLimitJsonObjectImpl{
		limitJsonObject: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
