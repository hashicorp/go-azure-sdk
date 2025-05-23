package recoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderSpecificRecoveryPointDetails interface {
	ProviderSpecificRecoveryPointDetails() BaseProviderSpecificRecoveryPointDetailsImpl
}

var _ ProviderSpecificRecoveryPointDetails = BaseProviderSpecificRecoveryPointDetailsImpl{}

type BaseProviderSpecificRecoveryPointDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseProviderSpecificRecoveryPointDetailsImpl) ProviderSpecificRecoveryPointDetails() BaseProviderSpecificRecoveryPointDetailsImpl {
	return s
}

var _ ProviderSpecificRecoveryPointDetails = RawProviderSpecificRecoveryPointDetailsImpl{}

// RawProviderSpecificRecoveryPointDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProviderSpecificRecoveryPointDetailsImpl struct {
	providerSpecificRecoveryPointDetails BaseProviderSpecificRecoveryPointDetailsImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawProviderSpecificRecoveryPointDetailsImpl) ProviderSpecificRecoveryPointDetails() BaseProviderSpecificRecoveryPointDetailsImpl {
	return s.providerSpecificRecoveryPointDetails
}

func UnmarshalProviderSpecificRecoveryPointDetailsImplementation(input []byte) (ProviderSpecificRecoveryPointDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProviderSpecificRecoveryPointDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2ARecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2ARecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageAzureV2") {
		var out InMageAzureV2RecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageAzureV2RecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageRcm") {
		var out InMageRcmRecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageRcmRecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseProviderSpecificRecoveryPointDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProviderSpecificRecoveryPointDetailsImpl: %+v", err)
	}

	return RawProviderSpecificRecoveryPointDetailsImpl{
		providerSpecificRecoveryPointDetails: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
