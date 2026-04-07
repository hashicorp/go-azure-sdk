package agentapplication

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationAuthorizationPolicy = OrganizationSharedBuiltInAuthorizationPolicy{}

type OrganizationSharedBuiltInAuthorizationPolicy struct {

	// Fields inherited from ApplicationAuthorizationPolicy

	Type BuiltInAuthorizationScheme `json:"type"`
}

func (s OrganizationSharedBuiltInAuthorizationPolicy) ApplicationAuthorizationPolicy() BaseApplicationAuthorizationPolicyImpl {
	return BaseApplicationAuthorizationPolicyImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = OrganizationSharedBuiltInAuthorizationPolicy{}

func (s OrganizationSharedBuiltInAuthorizationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper OrganizationSharedBuiltInAuthorizationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OrganizationSharedBuiltInAuthorizationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OrganizationSharedBuiltInAuthorizationPolicy: %+v", err)
	}

	decoded["type"] = "OrganizationScope"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OrganizationSharedBuiltInAuthorizationPolicy: %+v", err)
	}

	return encoded, nil
}
