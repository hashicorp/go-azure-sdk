package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterUnplannedFailoverProviderSpecificInput interface {
	ClusterUnplannedFailoverProviderSpecificInput() BaseClusterUnplannedFailoverProviderSpecificInputImpl
}

var _ ClusterUnplannedFailoverProviderSpecificInput = BaseClusterUnplannedFailoverProviderSpecificInputImpl{}

type BaseClusterUnplannedFailoverProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseClusterUnplannedFailoverProviderSpecificInputImpl) ClusterUnplannedFailoverProviderSpecificInput() BaseClusterUnplannedFailoverProviderSpecificInputImpl {
	return s
}

var _ ClusterUnplannedFailoverProviderSpecificInput = RawClusterUnplannedFailoverProviderSpecificInputImpl{}

// RawClusterUnplannedFailoverProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClusterUnplannedFailoverProviderSpecificInputImpl struct {
	clusterUnplannedFailoverProviderSpecificInput BaseClusterUnplannedFailoverProviderSpecificInputImpl
	Type                                          string
	Values                                        map[string]interface{}
}

func (s RawClusterUnplannedFailoverProviderSpecificInputImpl) ClusterUnplannedFailoverProviderSpecificInput() BaseClusterUnplannedFailoverProviderSpecificInputImpl {
	return s.clusterUnplannedFailoverProviderSpecificInput
}

func UnmarshalClusterUnplannedFailoverProviderSpecificInputImplementation(input []byte) (ClusterUnplannedFailoverProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterUnplannedFailoverProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AClusterUnplannedFailoverInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AClusterUnplannedFailoverInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseClusterUnplannedFailoverProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClusterUnplannedFailoverProviderSpecificInputImpl: %+v", err)
	}

	return RawClusterUnplannedFailoverProviderSpecificInputImpl{
		clusterUnplannedFailoverProviderSpecificInput: parent,
		Type:   value,
		Values: temp,
	}, nil

}
