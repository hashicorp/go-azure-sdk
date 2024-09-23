package actionrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionRuleProperties interface {
	ActionRuleProperties() BaseActionRulePropertiesImpl
}

var _ ActionRuleProperties = BaseActionRulePropertiesImpl{}

type BaseActionRulePropertiesImpl struct {
	Conditions     *Conditions       `json:"conditions,omitempty"`
	CreatedAt      *string           `json:"createdAt,omitempty"`
	CreatedBy      *string           `json:"createdBy,omitempty"`
	Description    *string           `json:"description,omitempty"`
	LastModifiedAt *string           `json:"lastModifiedAt,omitempty"`
	LastModifiedBy *string           `json:"lastModifiedBy,omitempty"`
	Scope          *Scope            `json:"scope,omitempty"`
	Status         *ActionRuleStatus `json:"status,omitempty"`
	Type           ActionRuleType    `json:"type"`
}

func (s BaseActionRulePropertiesImpl) ActionRuleProperties() BaseActionRulePropertiesImpl {
	return s
}

var _ ActionRuleProperties = RawActionRulePropertiesImpl{}

// RawActionRulePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawActionRulePropertiesImpl struct {
	actionRuleProperties BaseActionRulePropertiesImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawActionRulePropertiesImpl) ActionRuleProperties() BaseActionRulePropertiesImpl {
	return s.actionRuleProperties
}

func UnmarshalActionRulePropertiesImplementation(input []byte) (ActionRuleProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ActionRuleProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ActionGroup") {
		var out ActionGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ActionGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Diagnostics") {
		var out Diagnostics
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Diagnostics: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Suppression") {
		var out Suppression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Suppression: %+v", err)
		}
		return out, nil
	}

	var parent BaseActionRulePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseActionRulePropertiesImpl: %+v", err)
	}

	return RawActionRulePropertiesImpl{
		actionRuleProperties: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
