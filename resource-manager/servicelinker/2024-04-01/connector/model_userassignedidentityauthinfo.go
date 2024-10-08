package connector

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthInfoBase = UserAssignedIdentityAuthInfo{}

type UserAssignedIdentityAuthInfo struct {
	ClientId               *string                 `json:"clientId,omitempty"`
	DeleteOrUpdateBehavior *DeleteOrUpdateBehavior `json:"deleteOrUpdateBehavior,omitempty"`
	Roles                  *[]string               `json:"roles,omitempty"`
	SubscriptionId         *string                 `json:"subscriptionId,omitempty"`
	UserName               *string                 `json:"userName,omitempty"`

	// Fields inherited from AuthInfoBase

	AuthMode *AuthMode `json:"authMode,omitempty"`
	AuthType AuthType  `json:"authType"`
}

func (s UserAssignedIdentityAuthInfo) AuthInfoBase() BaseAuthInfoBaseImpl {
	return BaseAuthInfoBaseImpl{
		AuthMode: s.AuthMode,
		AuthType: s.AuthType,
	}
}

var _ json.Marshaler = UserAssignedIdentityAuthInfo{}

func (s UserAssignedIdentityAuthInfo) MarshalJSON() ([]byte, error) {
	type wrapper UserAssignedIdentityAuthInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserAssignedIdentityAuthInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserAssignedIdentityAuthInfo: %+v", err)
	}

	decoded["authType"] = "userAssignedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserAssignedIdentityAuthInfo: %+v", err)
	}

	return encoded, nil
}
