package channels

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceMoveChangeHistory struct {
	AzureSubscriptionId *string `json:"azureSubscriptionId,omitempty"`
	ChangedTimeUtc      *string `json:"changedTimeUtc,omitempty"`
	ResourceGroupName   *string `json:"resourceGroupName,omitempty"`
}

func (o *ResourceMoveChangeHistory) GetChangedTimeUtcAsTime() (*time.Time, error) {
	if o.ChangedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ChangedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ResourceMoveChangeHistory) SetChangedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ChangedTimeUtc = &formatted
}
