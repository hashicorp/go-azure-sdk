package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryAvailabilitySetCustomDetails interface {
}

func unmarshalRecoveryAvailabilitySetCustomDetailsImplementation(input []byte) (RecoveryAvailabilitySetCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryAvailabilitySetCustomDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["resourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingRecoveryAvailabilitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingRecoveryAvailabilitySet: %+v", err)
		}
		return out, nil
	}

	type RawRecoveryAvailabilitySetCustomDetailsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRecoveryAvailabilitySetCustomDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
