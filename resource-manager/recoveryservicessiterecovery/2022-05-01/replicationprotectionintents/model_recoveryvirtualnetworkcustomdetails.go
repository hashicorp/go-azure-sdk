package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryVirtualNetworkCustomDetails interface {
}

func unmarshalRecoveryVirtualNetworkCustomDetailsImplementation(input []byte) (RecoveryVirtualNetworkCustomDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RecoveryVirtualNetworkCustomDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["resourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Existing") {
		var out ExistingRecoveryVirtualNetwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExistingRecoveryVirtualNetwork: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "New") {
		var out NewRecoveryVirtualNetwork
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NewRecoveryVirtualNetwork: %+v", err)
		}
		return out, nil
	}

	type RawRecoveryVirtualNetworkCustomDetailsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRecoveryVirtualNetworkCustomDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
