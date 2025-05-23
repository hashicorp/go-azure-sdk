package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateProtectionIntentProviderSpecificDetails interface {
	CreateProtectionIntentProviderSpecificDetails() BaseCreateProtectionIntentProviderSpecificDetailsImpl
}

var _ CreateProtectionIntentProviderSpecificDetails = BaseCreateProtectionIntentProviderSpecificDetailsImpl{}

type BaseCreateProtectionIntentProviderSpecificDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseCreateProtectionIntentProviderSpecificDetailsImpl) CreateProtectionIntentProviderSpecificDetails() BaseCreateProtectionIntentProviderSpecificDetailsImpl {
	return s
}

var _ CreateProtectionIntentProviderSpecificDetails = RawCreateProtectionIntentProviderSpecificDetailsImpl{}

// RawCreateProtectionIntentProviderSpecificDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCreateProtectionIntentProviderSpecificDetailsImpl struct {
	createProtectionIntentProviderSpecificDetails BaseCreateProtectionIntentProviderSpecificDetailsImpl
	Type                                          string
	Values                                        map[string]interface{}
}

func (s RawCreateProtectionIntentProviderSpecificDetailsImpl) CreateProtectionIntentProviderSpecificDetails() BaseCreateProtectionIntentProviderSpecificDetailsImpl {
	return s.createProtectionIntentProviderSpecificDetails
}

func UnmarshalCreateProtectionIntentProviderSpecificDetailsImplementation(input []byte) (CreateProtectionIntentProviderSpecificDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CreateProtectionIntentProviderSpecificDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2ACreateProtectionIntentInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2ACreateProtectionIntentInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseCreateProtectionIntentProviderSpecificDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCreateProtectionIntentProviderSpecificDetailsImpl: %+v", err)
	}

	return RawCreateProtectionIntentProviderSpecificDetailsImpl{
		createProtectionIntentProviderSpecificDetails: parent,
		Type:   value,
		Values: temp,
	}, nil

}
