package replicationappliances

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DraDetails struct {
	BiosId                    *string           `json:"biosId,omitempty"`
	ForwardProtectedItemCount *int64            `json:"forwardProtectedItemCount,omitempty"`
	Health                    *ProtectionHealth `json:"health,omitempty"`
	HealthErrors              *[]HealthError    `json:"healthErrors,omitempty"`
	Id                        *string           `json:"id,omitempty"`
	LastHeartbeatUtc          *string           `json:"lastHeartbeatUtc,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	ReverseProtectedItemCount *int64            `json:"reverseProtectedItemCount,omitempty"`
	Version                   *string           `json:"version,omitempty"`
}

func (o *DraDetails) GetLastHeartbeatUtcAsTime() (*time.Time, error) {
	if o.LastHeartbeatUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastHeartbeatUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *DraDetails) SetLastHeartbeatUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastHeartbeatUtc = &formatted
}
