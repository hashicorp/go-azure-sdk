package usersession

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSessionProperties struct {
	ActiveDirectoryUserName *string          `json:"activeDirectoryUserName,omitempty"`
	ApplicationType         *ApplicationType `json:"applicationType,omitempty"`
	CreateTime              *string          `json:"createTime,omitempty"`
	ObjectId                *string          `json:"objectId,omitempty"`
	SessionState            *SessionState    `json:"sessionState,omitempty"`
	UserPrincipalName       *string          `json:"userPrincipalName,omitempty"`
}

func (o *UserSessionProperties) GetCreateTimeAsTime() (*time.Time, error) {
	if o.CreateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UserSessionProperties) SetCreateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreateTime = &formatted
}
