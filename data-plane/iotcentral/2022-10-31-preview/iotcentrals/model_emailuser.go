package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ User = EmailUser{}

type EmailUser struct {
	Email string `json:"email"`

	// Fields inherited from User

	Id    *string          `json:"id,omitempty"`
	Roles []RoleAssignment `json:"roles"`
	Type  string           `json:"type"`
}

func (s EmailUser) User() BaseUserImpl {
	return BaseUserImpl{
		Id:    s.Id,
		Roles: s.Roles,
		Type:  s.Type,
	}
}

var _ json.Marshaler = EmailUser{}

func (s EmailUser) MarshalJSON() ([]byte, error) {
	type wrapper EmailUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmailUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmailUser: %+v", err)
	}

	decoded["type"] = "email"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmailUser: %+v", err)
	}

	return encoded, nil
}
