package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EarlyTerminationPolicy interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEarlyTerminationPolicyImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalEarlyTerminationPolicyImplementation(input []byte) (EarlyTerminationPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EarlyTerminationPolicy into map[string]interface: %+v", err)
	}

	value, ok := temp["policyType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bandit") {
		var out BanditPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BanditPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MedianStopping") {
		var out MedianStoppingPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MedianStoppingPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TruncationSelection") {
		var out TruncationSelectionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TruncationSelectionPolicy: %+v", err)
		}
		return out, nil
	}

	out := RawEarlyTerminationPolicyImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
