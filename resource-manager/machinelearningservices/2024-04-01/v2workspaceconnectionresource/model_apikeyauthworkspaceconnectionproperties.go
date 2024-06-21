package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = ApiKeyAuthWorkspaceConnectionProperties{}

type ApiKeyAuthWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionApiKey `json:"credentials,omitempty"`

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

var _ json.Marshaler = ApiKeyAuthWorkspaceConnectionProperties{}

func (s ApiKeyAuthWorkspaceConnectionProperties) MarshalJSON() ([]byte, error) {
	type wrapper ApiKeyAuthWorkspaceConnectionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApiKeyAuthWorkspaceConnectionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApiKeyAuthWorkspaceConnectionProperties: %+v", err)
	}
	decoded["authType"] = "ApiKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApiKeyAuthWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
