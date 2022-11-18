package recoverypoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderSpecificRecoveryPointDetails interface {
}

func unmarshalProviderSpecificRecoveryPointDetailsImplementation(input []byte) (ProviderSpecificRecoveryPointDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProviderSpecificRecoveryPointDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out A2ARecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2ARecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageAzureV2") {
		var out InMageAzureV2RecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageAzureV2RecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InMageRcm") {
		var out InMageRcmRecoveryPointDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InMageRcmRecoveryPointDetails: %+v", err)
		}
		return out, nil
	}

	type RawProviderSpecificRecoveryPointDetailsImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawProviderSpecificRecoveryPointDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
