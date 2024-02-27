package vcenters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceStatus struct {
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
	Message       *string `json:"message,omitempty"`
	Reason        *string `json:"reason,omitempty"`
	Severity      *string `json:"severity,omitempty"`
	Status        *string `json:"status,omitempty"`
	Type          *string `json:"type,omitempty"`
}

func (o *ResourceStatus) GetLastUpdatedAtAsTime() (*time.Time, error) {
	if o.LastUpdatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedAt, "2006-01-02T15:04:05Z07:00")
}
