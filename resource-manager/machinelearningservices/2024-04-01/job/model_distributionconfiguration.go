package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributionConfiguration interface {
	DistributionConfiguration() BaseDistributionConfigurationImpl
}

var _ DistributionConfiguration = BaseDistributionConfigurationImpl{}

type BaseDistributionConfigurationImpl struct {
	DistributionType DistributionType `json:"distributionType"`
}

func (s BaseDistributionConfigurationImpl) DistributionConfiguration() BaseDistributionConfigurationImpl {
	return s
}

var _ DistributionConfiguration = RawDistributionConfigurationImpl{}

// RawDistributionConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDistributionConfigurationImpl struct {
	distributionConfiguration BaseDistributionConfigurationImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawDistributionConfigurationImpl) DistributionConfiguration() BaseDistributionConfigurationImpl {
	return s.distributionConfiguration
}

func UnmarshalDistributionConfigurationImplementation(input []byte) (DistributionConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DistributionConfiguration into map[string]interface: %+v", err)
	}

	value, ok := temp["distributionType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Mpi") {
		var out Mpi
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Mpi: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PyTorch") {
		var out PyTorch
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PyTorch: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TensorFlow") {
		var out TensorFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TensorFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseDistributionConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDistributionConfigurationImpl: %+v", err)
	}

	return RawDistributionConfigurationImpl{
		distributionConfiguration: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
