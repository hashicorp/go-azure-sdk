package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = SASAuthTypeWorkspaceConnectionProperties{}

type SASAuthTypeWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionSharedAccessSignature `json:"credentials,omitempty"`

	// Fields inherited from WorkspaceConnectionPropertiesV2

	AuthType                    ConnectionAuthType    `json:"authType"`
	Category                    *ConnectionCategory   `json:"category,omitempty"`
	CreatedByWorkspaceArmId     *string               `json:"createdByWorkspaceArmId,omitempty"`
	Error                       *string               `json:"error,omitempty"`
	ExpiryTime                  *string               `json:"expiryTime,omitempty"`
	Group                       *ConnectionGroup      `json:"group,omitempty"`
	IsSharedToAll               *bool                 `json:"isSharedToAll,omitempty"`
	Metadata                    *map[string]string    `json:"metadata,omitempty"`
	PeRequirement               *ManagedPERequirement `json:"peRequirement,omitempty"`
	PeStatus                    *ManagedPEStatus      `json:"peStatus,omitempty"`
	SharedUserList              *[]string             `json:"sharedUserList,omitempty"`
	Target                      *string               `json:"target,omitempty"`
	UseWorkspaceManagedIdentity *bool                 `json:"useWorkspaceManagedIdentity,omitempty"`
}

func (s SASAuthTypeWorkspaceConnectionProperties) WorkspaceConnectionPropertiesV2() BaseWorkspaceConnectionPropertiesV2Impl {
	return BaseWorkspaceConnectionPropertiesV2Impl{
		AuthType:                    s.AuthType,
		Category:                    s.Category,
		CreatedByWorkspaceArmId:     s.CreatedByWorkspaceArmId,
		Error:                       s.Error,
		ExpiryTime:                  s.ExpiryTime,
		Group:                       s.Group,
		IsSharedToAll:               s.IsSharedToAll,
		Metadata:                    s.Metadata,
		PeRequirement:               s.PeRequirement,
		PeStatus:                    s.PeStatus,
		SharedUserList:              s.SharedUserList,
		Target:                      s.Target,
		UseWorkspaceManagedIdentity: s.UseWorkspaceManagedIdentity,
	}
}

func (o *SASAuthTypeWorkspaceConnectionProperties) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SASAuthTypeWorkspaceConnectionProperties) SetExpiryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTime = &formatted
}

var _ json.Marshaler = SASAuthTypeWorkspaceConnectionProperties{}

func (s SASAuthTypeWorkspaceConnectionProperties) MarshalJSON() ([]byte, error) {
	type wrapper SASAuthTypeWorkspaceConnectionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SASAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SASAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	decoded["authType"] = "SAS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SASAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
