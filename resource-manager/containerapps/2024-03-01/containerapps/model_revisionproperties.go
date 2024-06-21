package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevisionProperties struct {
	Active            *bool                      `json:"active,omitempty"`
	CreatedTime       *string                    `json:"createdTime,omitempty"`
	Fqdn              *string                    `json:"fqdn,omitempty"`
	HealthState       *RevisionHealthState       `json:"healthState,omitempty"`
	LastActiveTime    *string                    `json:"lastActiveTime,omitempty"`
	ProvisioningError *string                    `json:"provisioningError,omitempty"`
	ProvisioningState *RevisionProvisioningState `json:"provisioningState,omitempty"`
	Replicas          *int64                     `json:"replicas,omitempty"`
	RunningState      *RevisionRunningState      `json:"runningState,omitempty"`
	Template          *Template                  `json:"template,omitempty"`
	TrafficWeight     *int64                     `json:"trafficWeight,omitempty"`
}
