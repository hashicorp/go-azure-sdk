package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ User = ADGroupUser{}

type ADGroupUser struct {
	ObjectId string `json:"objectId"`
	TenantId string `json:"tenantId"`

	// Fields inherited from User

	Id    *string          `json:"id,omitempty"`
	Roles []RoleAssignment `json:"roles"`
	Type  string           `json:"type"`
}

func (s ADGroupUser) User() BaseUserImpl {
	return BaseUserImpl{
		Id:    s.Id,
		Roles: s.Roles,
		Type:  s.Type,
	}
}

var _ json.Marshaler = ADGroupUser{}

func (s ADGroupUser) MarshalJSON() ([]byte, error) {
	type wrapper ADGroupUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ADGroupUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ADGroupUser: %+v", err)
	}

	decoded["type"] = "adGroup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ADGroupUser: %+v", err)
	}

	return encoded, nil
}
