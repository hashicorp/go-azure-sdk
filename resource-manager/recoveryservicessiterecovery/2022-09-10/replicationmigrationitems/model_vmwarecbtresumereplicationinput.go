package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResumeReplicationProviderSpecificInput = VMwareCbtResumeReplicationInput{}

type VMwareCbtResumeReplicationInput struct {
	DeleteMigrationResources *string `json:"deleteMigrationResources,omitempty"`

	// Fields inherited from ResumeReplicationProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtResumeReplicationInput{}

func (s VMwareCbtResumeReplicationInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtResumeReplicationInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtResumeReplicationInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtResumeReplicationInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtResumeReplicationInput: %+v", err)
	}

	return encoded, nil
}
