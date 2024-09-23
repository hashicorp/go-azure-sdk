package contentkeypolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyConfiguration interface {
	ContentKeyPolicyConfiguration() BaseContentKeyPolicyConfigurationImpl
}

var _ ContentKeyPolicyConfiguration = BaseContentKeyPolicyConfigurationImpl{}

type BaseContentKeyPolicyConfigurationImpl struct {
	OdataType string `json:"@odata.type"`
}

func (s BaseContentKeyPolicyConfigurationImpl) ContentKeyPolicyConfiguration() BaseContentKeyPolicyConfigurationImpl {
	return s
}

var _ ContentKeyPolicyConfiguration = RawContentKeyPolicyConfigurationImpl{}

// RawContentKeyPolicyConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawContentKeyPolicyConfigurationImpl struct {
	contentKeyPolicyConfiguration BaseContentKeyPolicyConfigurationImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawContentKeyPolicyConfigurationImpl) ContentKeyPolicyConfiguration() BaseContentKeyPolicyConfigurationImpl {
	return s.contentKeyPolicyConfiguration
}

func UnmarshalContentKeyPolicyConfigurationImplementation(input []byte) (ContentKeyPolicyConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ContentKeyPolicyConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration") {
		var out ContentKeyPolicyClearKeyConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyClearKeyConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyFairPlayConfiguration") {
		var out ContentKeyPolicyFairPlayConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyFairPlayConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration") {
		var out ContentKeyPolicyPlayReadyConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyPlayReadyConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyUnknownConfiguration") {
		var out ContentKeyPolicyUnknownConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyUnknownConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration") {
		var out ContentKeyPolicyWidevineConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContentKeyPolicyWidevineConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseContentKeyPolicyConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseContentKeyPolicyConfigurationImpl: %+v", err)
	}

	return RawContentKeyPolicyConfigurationImpl{
		contentKeyPolicyConfiguration: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
