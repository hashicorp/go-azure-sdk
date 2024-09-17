package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionProfileCustomDetails interface {
	ProtectionProfileCustomDetails() BaseProtectionProfileCustomDetailsImpl
}

var _ ProtectionProfileCustomDetails = BaseProtectionProfileCustomDetailsImpl{}

type BaseProtectionProfileCustomDetailsImpl struct {
	ResourceType string `json:"resourceType"`
}

func (s BaseProtectionProfileCustomDetailsImpl) ProtectionProfileCustomDetails() BaseProtectionProfileCustomDetailsImpl {
	return s
}

var _ ProtectionProfileCustomDetails = RawProtectionProfileCustomDetailsImpl{}

// RawProtectionProfileCustomDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionProfileCustomDetailsImpl struct {
	protectionProfileCustomDetails BaseProtectionProfileCustomDetailsImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawProtectionProfileCustomDetailsImpl) ProtectionProfileCustomDetails() BaseProtectionProfileCustomDetailsImpl {
	return s.protectionProfileCustomDetails
}

func UnmarshalProtectionProfileCustomDetailsImplementation(input []byte) (ProtectionProfileCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionProfileCustomDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["resourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingProtectionProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingProtectionProfile: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "New") {
		var out NewProtectionProfile
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NewProtectionProfile: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionProfileCustomDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionProfileCustomDetailsImpl: %+v", err)
	}

	return RawProtectionProfileCustomDetailsImpl{
		protectionProfileCustomDetails: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
