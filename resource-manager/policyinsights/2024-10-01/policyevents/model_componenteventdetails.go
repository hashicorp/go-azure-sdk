package policyevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentEventDetails struct {
	Id                     *string `json:"id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	PrincipalOid           *string `json:"principalOid,omitempty"`
	TenantId               *string `json:"tenantId,omitempty"`
	Timestamp              *string `json:"timestamp,omitempty"`
	Type                   *string `json:"type,omitempty"`
}

func (o *ComponentEventDetails) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *ComponentEventDetails) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
