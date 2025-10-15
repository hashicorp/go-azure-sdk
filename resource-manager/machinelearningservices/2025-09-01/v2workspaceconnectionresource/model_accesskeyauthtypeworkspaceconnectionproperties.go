package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WorkspaceConnectionPropertiesV2 = AccessKeyAuthTypeWorkspaceConnectionProperties{}

type AccessKeyAuthTypeWorkspaceConnectionProperties struct {
	Credentials *WorkspaceConnectionAccessKey `json:"credentials,omitempty"`

	// Fields inherited from WorkspaceConnectionPropertiesV2

	AuthType                ConnectionAuthType  `json:"authType"`
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

func (s AccessKeyAuthTypeWorkspaceConnectionProperties) WorkspaceConnectionPropertiesV2() BaseWorkspaceConnectionPropertiesV2Impl {
	return BaseWorkspaceConnectionPropertiesV2Impl{
		AuthType:                s.AuthType,
		Category:                s.Category,
		CreatedByWorkspaceArmId: s.CreatedByWorkspaceArmId,
		ExpiryTime:              s.ExpiryTime,
		Group:                   s.Group,
		IsSharedToAll:           s.IsSharedToAll,
		Metadata:                s.Metadata,
		SharedUserList:          s.SharedUserList,
		Target:                  s.Target,
		Value:                   s.Value,
		ValueFormat:             s.ValueFormat,
	}
}

func (o *AccessKeyAuthTypeWorkspaceConnectionProperties) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AccessKeyAuthTypeWorkspaceConnectionProperties) SetExpiryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTime = &formatted
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	decoded["authType"] = "AccessKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessKeyAuthTypeWorkspaceConnectionProperties: %+v", err)
	}

	return encoded, nil
}
