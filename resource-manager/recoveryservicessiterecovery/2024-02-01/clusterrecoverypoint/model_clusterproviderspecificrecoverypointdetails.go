package clusterrecoverypoint

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProviderSpecificRecoveryPointDetails interface {
}

// RawClusterProviderSpecificRecoveryPointDetailsImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClusterProviderSpecificRecoveryPointDetailsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalClusterProviderSpecificRecoveryPointDetailsImplementation(input []byte) (ClusterProviderSpecificRecoveryPointDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterProviderSpecificRecoveryPointDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AClusterRecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AClusterRecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	out := RawClusterProviderSpecificRecoveryPointDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
