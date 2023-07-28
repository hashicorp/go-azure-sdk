package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = AccessKeyAuthTypeWorkspaceConnectionProperties{}

type AccessKeyAuthTypeWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionAccessKey `json:"credentials,omitempty"`

	// Fields inherited from WorkspaceConnectionPropertiesV2
	Category    *ConnectionCategory `json:"category,omitempty"`
	ExpiryTime  *string             `json:"expiryTime,omitempty"`
	Target      *string             `json:"target,omitempty"`
	Value       *string             `json:"value,omitempty"`
	ValueFormat *ValueFormat        `json:"valueFormat,omitempty"`
}

var _ json.Marshaler = AccessKeyAuthTypeWorkspaceConnectionProperties{}

func (s AccessKeyAuthTypeWorkspaceConnectionProperties) MarshalJSON() ([]byte, error) {
	type wrapper AccessKeyAuthTypeWorkspaceConnectionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}
	decoded["authType"] = "AccessKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
