package featuresupport

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureSupportRequest interface {
	FeatureSupportRequest() BaseFeatureSupportRequestImpl
}

var _ FeatureSupportRequest = BaseFeatureSupportRequestImpl{}

type BaseFeatureSupportRequestImpl struct {
	FeatureType string `json:"featureType"`
}

func (s BaseFeatureSupportRequestImpl) FeatureSupportRequest() BaseFeatureSupportRequestImpl {
	return s
}

var _ FeatureSupportRequest = RawFeatureSupportRequestImpl{}

// RawFeatureSupportRequestImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFeatureSupportRequestImpl struct {
	featureSupportRequest BaseFeatureSupportRequestImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawFeatureSupportRequestImpl) FeatureSupportRequest() BaseFeatureSupportRequestImpl {
	return s.featureSupportRequest
}

func UnmarshalFeatureSupportRequestImplementation(input []byte) (FeatureSupportRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureSupportRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["featureType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureBackupGoals") {
		var out AzureBackupGoalFeatureSupportRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupGoalFeatureSupportRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureVMResourceBackup") {
		var out AzureVMResourceFeatureSupportRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureVMResourceFeatureSupportRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseFeatureSupportRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFeatureSupportRequestImpl: %+v", err)
	}

	return RawFeatureSupportRequestImpl{
		featureSupportRequest: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
