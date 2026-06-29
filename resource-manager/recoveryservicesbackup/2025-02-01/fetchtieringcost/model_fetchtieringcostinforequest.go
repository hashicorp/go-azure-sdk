package fetchtieringcost

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FetchTieringCostInfoRequest interface {
	FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl
}

var _ FetchTieringCostInfoRequest = BaseFetchTieringCostInfoRequestImpl{}

type BaseFetchTieringCostInfoRequestImpl struct {
	ObjectType     string                `json:"objectType"`
	SourceTierType RecoveryPointTierType `json:"sourceTierType"`
	TargetTierType RecoveryPointTierType `json:"targetTierType"`
}

func (s BaseFetchTieringCostInfoRequestImpl) FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl {
	return s
}

var _ FetchTieringCostInfoRequest = RawFetchTieringCostInfoRequestImpl{}

// RawFetchTieringCostInfoRequestImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawFetchTieringCostInfoRequestImpl struct {
	fetchTieringCostInfoRequest BaseFetchTieringCostInfoRequestImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawFetchTieringCostInfoRequestImpl) FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl {
	return s.fetchTieringCostInfoRequest
}

func (s RawFetchTieringCostInfoRequestImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalFetchTieringCostInfoRequestImplementation(input []byte) (FetchTieringCostInfoRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FetchTieringCostInfoRequest into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "FetchTieringCostInfoForRehydrationRequest") {
		var out FetchTieringCostInfoForRehydrationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FetchTieringCostInfoForRehydrationRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FetchTieringCostSavingsInfoForPolicyRequest") {
		var out FetchTieringCostSavingsInfoForPolicyRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FetchTieringCostSavingsInfoForPolicyRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FetchTieringCostSavingsInfoForProtectedItemRequest") {
		var out FetchTieringCostSavingsInfoForProtectedItemRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FetchTieringCostSavingsInfoForProtectedItemRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FetchTieringCostSavingsInfoForVaultRequest") {
		var out FetchTieringCostSavingsInfoForVaultRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FetchTieringCostSavingsInfoForVaultRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseFetchTieringCostInfoRequestImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFetchTieringCostInfoRequestImpl: %+v", err)
	}

	return RawFetchTieringCostInfoRequestImpl{
		fetchTieringCostInfoRequest: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
