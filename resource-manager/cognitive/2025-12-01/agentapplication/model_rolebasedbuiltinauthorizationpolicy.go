package agentapplication

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationAuthorizationPolicy = RoleBasedBuiltInAuthorizationPolicy{}

type RoleBasedBuiltInAuthorizationPolicy struct {

	// Fields inherited from ApplicationAuthorizationPolicy

	Type BuiltInAuthorizationScheme `json:"type"`
}

func (s RoleBasedBuiltInAuthorizationPolicy) ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl {
	return BaseApplicationAuthorizationPolicyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = RoleBasedBuiltInAuthorizationPolicy{}

func (s RoleBasedBuiltInAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper RoleBasedBuiltInAuthorizationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RoleBasedBuiltInAuthorizationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RoleBasedBuiltInAuthorizationPolicy: %+v", err)
	}

	decoded["type"] = "Default"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RoleBasedBuiltInAuthorizationPolicy: %+v", err)
	}

	return encoded, nil
}
