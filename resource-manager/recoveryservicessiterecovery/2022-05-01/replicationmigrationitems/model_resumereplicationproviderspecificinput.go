package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResumeReplicationProviderSpecificInput interface {
}

func unmarshalResumeReplicationProviderSpecificInputImplementation(input []byte) (ResumeReplicationProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResumeReplicationProviderSpecificInput into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtResumeReplicationInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtResumeReplicationInput: %+v", err)
		}
		return out, nil
	}

	type RawResumeReplicationProviderSpecificInputImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawResumeReplicationProviderSpecificInputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
