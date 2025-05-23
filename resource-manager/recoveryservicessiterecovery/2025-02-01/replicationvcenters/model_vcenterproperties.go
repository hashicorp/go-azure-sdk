package replicationvcenters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterProperties struct {
	DiscoveryStatus       *string        `json:"discoveryStatus,omitempty"`
	FabricArmResourceName *string        `json:"fabricArmResourceName,omitempty"`
	FriendlyName          *string        `json:"friendlyName,omitempty"`
	HealthErrors          *[]HealthError `json:"healthErrors,omitempty"`
	IPAddress             *string        `json:"ipAddress,omitempty"`
	InfrastructureId      *string        `json:"infrastructureId,omitempty"`
	InternalId            *string        `json:"internalId,omitempty"`
	LastHeartbeat         *string        `json:"lastHeartbeat,omitempty"`
	Port                  *string        `json:"port,omitempty"`
	ProcessServerId       *string        `json:"processServerId,omitempty"`
	RunAsAccountId        *string        `json:"runAsAccountId,omitempty"`
}

func (o *VCenterProperties) GetLastHeartbeatAsTime() (*time.Time, error) {
	if o.LastHeartbeat == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastHeartbeat, "2006-01-02T15:04:05Z07:00")
}

func (o *VCenterProperties) SetLastHeartbeatAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastHeartbeat = &formatted
}
