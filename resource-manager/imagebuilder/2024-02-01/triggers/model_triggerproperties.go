package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerProperties interface {
	TriggerProperties() BaseTriggerPropertiesImpl
}

var _ TriggerProperties = BaseTriggerPropertiesImpl{}

type BaseTriggerPropertiesImpl struct {
	Kind              string             `json:"kind"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	Status            *TriggerStatus     `json:"status,omitempty"`
}

func (s BaseTriggerPropertiesImpl) TriggerProperties() BaseTriggerPropertiesImpl {
	return s
}

var _ TriggerProperties = RawTriggerPropertiesImpl{}

// RawTriggerPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTriggerPropertiesImpl struct {
	triggerProperties BaseTriggerPropertiesImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawTriggerPropertiesImpl) TriggerProperties() BaseTriggerPropertiesImpl {
	return s.triggerProperties
}

func UnmarshalTriggerPropertiesImplementation(input []byte) (TriggerProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TriggerProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "SourceImage") {
		var out SourceImageTriggerProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SourceImageTriggerProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseTriggerPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTriggerPropertiesImpl: %+v", err)
	}

	return RawTriggerPropertiesImpl{
		triggerProperties: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
