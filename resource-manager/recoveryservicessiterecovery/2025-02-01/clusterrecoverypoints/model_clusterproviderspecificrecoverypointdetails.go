package clusterrecoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProviderSpecificRecoveryPointDetails interface {
	ClusterProviderSpecificRecoveryPointDetails() BaseClusterProviderSpecificRecoveryPointDetailsImpl
}

var _ ClusterProviderSpecificRecoveryPointDetails = BaseClusterProviderSpecificRecoveryPointDetailsImpl{}

type BaseClusterProviderSpecificRecoveryPointDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseClusterProviderSpecificRecoveryPointDetailsImpl) ClusterProviderSpecificRecoveryPointDetails() BaseClusterProviderSpecificRecoveryPointDetailsImpl {
	return s
}

var _ ClusterProviderSpecificRecoveryPointDetails = RawClusterProviderSpecificRecoveryPointDetailsImpl{}

// RawClusterProviderSpecificRecoveryPointDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawClusterProviderSpecificRecoveryPointDetailsImpl struct {
	clusterProviderSpecificRecoveryPointDetails BaseClusterProviderSpecificRecoveryPointDetailsImpl
	Type                                        string
	Values                                      map[string]interface{}
}

func (s RawClusterProviderSpecificRecoveryPointDetailsImpl) ClusterProviderSpecificRecoveryPointDetails() BaseClusterProviderSpecificRecoveryPointDetailsImpl {
	return s.clusterProviderSpecificRecoveryPointDetails
}

func (s RawClusterProviderSpecificRecoveryPointDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalClusterProviderSpecificRecoveryPointDetailsImplementation(input []byte) (ClusterProviderSpecificRecoveryPointDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterProviderSpecificRecoveryPointDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AClusterRecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AClusterRecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseClusterProviderSpecificRecoveryPointDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClusterProviderSpecificRecoveryPointDetailsImpl: %+v", err)
	}

	return RawClusterProviderSpecificRecoveryPointDetailsImpl{
		clusterProviderSpecificRecoveryPointDetails: parent,
		Type:   value,
		Values: temp,
	}, nil

}
