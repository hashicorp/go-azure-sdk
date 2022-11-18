package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TestMigrateProviderSpecificInput = VMwareCbtTestMigrateInput{}

type VMwareCbtTestMigrateInput struct {
	NetworkId       string               `json:"networkId"`
	RecoveryPointId string               `json:"recoveryPointId"`
	VmNics          *[]VMwareCbtNicInput `json:"vmNics,omitempty"`

	// Fields inherited from TestMigrateProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtTestMigrateInput{}

func (s VMwareCbtTestMigrateInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtTestMigrateInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtTestMigrateInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtTestMigrateInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtTestMigrateInput: %+v", err)
	}

	return encoded, nil
}
