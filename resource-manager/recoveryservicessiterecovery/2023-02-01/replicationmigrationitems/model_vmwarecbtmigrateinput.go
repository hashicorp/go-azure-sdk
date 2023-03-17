package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateProviderSpecificInput = VMwareCbtMigrateInput{}

type VMwareCbtMigrateInput struct {
	OsUpgradeVersion *string `json:"osUpgradeVersion,omitempty"`
	PerformShutdown  string  `json:"performShutdown"`

	// Fields inherited from MigrateProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtMigrateInput{}

func (s VMwareCbtMigrateInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtMigrateInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtMigrateInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtMigrateInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtMigrateInput: %+v", err)
	}

	return encoded, nil
}
