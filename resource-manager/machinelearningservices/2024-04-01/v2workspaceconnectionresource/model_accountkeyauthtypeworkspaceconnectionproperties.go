package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = AccountKeyAuthTypeWorkspaceConnectionProperties{}

type AccountKeyAuthTypeWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionAccountKey `json:"credentials,omitempty"`

	// Fields inherited from WorkspaceConnectionPropertiesV2
	Category                *ConnectionCategory `json:"category,omitempty"`
	CreatedByWorkspaceArmId *string             `json:"createdByWorkspaceArmId,omitempty"`
	ExpiryTime              *string             `json:"expiryTime,omitempty"`
	Group                   *ConnectionGroup    `json:"group,omitempty"`
	IsSharedToAll           *bool               `json:"isSharedToAll,omitempty"`
	Metadata                *map[string]string  `json:"metadata,omitempty"`
	SharedUserList          *[]string           `json:"sharedUserList,omitempty"`
	Target                  *string             `json:"target,omitempty"`
	Value                   *string             `json:"value,omitempty"`
	ValueFormat             *ValueFormat        `json:"valueFormat,omitempty"`
}

var _ json.Marshaler = AccountKeyAuthTypeWorkspaceConnectionProperties{}

func (s AccountKeyAuthTypeWorkspaceConnectionProperties) MarshalJSON() ([]byte, error) {
	type wrapper AccountKeyAuthTypeWorkspaceConnectionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccountKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccountKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}
	decoded["authType"] = "AccountKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccountKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
