package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessEntityProperties struct {
	AccountEntityId          *string                 `json:"accountEntityId,omitempty"`
	AdditionalData           *map[string]interface{} `json:"additionalData,omitempty"`
	CommandLine              *string                 `json:"commandLine,omitempty"`
	CreationTimeUtc          *string                 `json:"creationTimeUtc,omitempty"`
	ElevationToken           *ElevationToken         `json:"elevationToken,omitempty"`
	FriendlyName             *string                 `json:"friendlyName,omitempty"`
	HostEntityId             *string                 `json:"hostEntityId,omitempty"`
	HostLogonSessionEntityId *string                 `json:"hostLogonSessionEntityId,omitempty"`
	ImageFileEntityId        *string                 `json:"imageFileEntityId,omitempty"`
	ParentProcessEntityId    *string                 `json:"parentProcessEntityId,omitempty"`
	ProcessId                *string                 `json:"processId,omitempty"`
}

func (o *ProcessEntityProperties) GetCreationTimeUtcAsTime() (*time.Time, error) {
	if o.CreationTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ProcessEntityProperties) SetCreationTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTimeUtc = &formatted
}
