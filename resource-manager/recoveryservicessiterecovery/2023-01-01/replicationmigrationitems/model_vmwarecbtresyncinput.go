package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ResyncProviderSpecificInput = VMwareCbtResyncInput{}

type VMwareCbtResyncInput struct {
	SkipCbtReset string `json:"skipCbtReset"`

	// Fields inherited from ResyncProviderSpecificInput
}

var _ json.Marshaler = VMwareCbtResyncInput{}

func (s VMwareCbtResyncInput) MarshalJSON() ([]byte, error) {
	type wrapper VMwareCbtResyncInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VMwareCbtResyncInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VMwareCbtResyncInput: %+v", err)
	}
	decoded["instanceType"] = "VMwareCbt"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VMwareCbtResyncInput: %+v", err)
	}

	return encoded, nil
}
