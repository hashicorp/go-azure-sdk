package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterTestFailoverProviderSpecificInput interface {
	ClusterTestFailoverProviderSpecificInput() BaseClusterTestFailoverProviderSpecificInputImpl
}

var _ ClusterTestFailoverProviderSpecificInput = BaseClusterTestFailoverProviderSpecificInputImpl{}

type BaseClusterTestFailoverProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseClusterTestFailoverProviderSpecificInputImpl) ClusterTestFailoverProviderSpecificInput() BaseClusterTestFailoverProviderSpecificInputImpl {
	return s
}

var _ ClusterTestFailoverProviderSpecificInput = RawClusterTestFailoverProviderSpecificInputImpl{}

// RawClusterTestFailoverProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClusterTestFailoverProviderSpecificInputImpl struct {
	clusterTestFailoverProviderSpecificInput BaseClusterTestFailoverProviderSpecificInputImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawClusterTestFailoverProviderSpecificInputImpl) ClusterTestFailoverProviderSpecificInput() BaseClusterTestFailoverProviderSpecificInputImpl {
	return s.clusterTestFailoverProviderSpecificInput
}

func UnmarshalClusterTestFailoverProviderSpecificInputImplementation(input []byte) (ClusterTestFailoverProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterTestFailoverProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AClusterTestFailoverInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AClusterTestFailoverInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseClusterTestFailoverProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClusterTestFailoverProviderSpecificInputImpl: %+v", err)
	}

	return RawClusterTestFailoverProviderSpecificInputImpl{
		clusterTestFailoverProviderSpecificInput: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
