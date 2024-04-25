package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = ManagedIdentityAuthTypeWorkspaceConnectionProperties{}

type ManagedIdentityAuthTypeWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionManagedIdentity `json:"credentials,omitempty"`

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

func (o *ManagedIdentityAuthTypeWorkspaceConnectionProperties) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedIdentityAuthTypeWorkspaceConnectionProperties) SetExpiryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTime = &formatted
}

var _ json.Marshaler = ManagedIdentityAuthTypeWorkspaceConnectionProperties{}

func (s ManagedIdentityAuthTypeWorkspaceConnectionProperties) MarshalJSON() ([]byte, error) {
	type wrapper ManagedIdentityAuthTypeWorkspaceConnectionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedIdentityAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedIdentityAuthTypeWorkspaceConnectionProperties: %+v", err)
	}
	decoded["authType"] = "ManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedIdentityAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
