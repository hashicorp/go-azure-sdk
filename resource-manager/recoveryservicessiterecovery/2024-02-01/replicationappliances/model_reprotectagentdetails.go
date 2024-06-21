package replicationappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReprotectAgentDetails struct {
	AccessibleDatastores *[]string         `json:"accessibleDatastores,omitempty"`
	BiosId               *string           `json:"biosId,omitempty"`
	FabricObjectId       *string           `json:"fabricObjectId,omitempty"`
	Fqdn                 *string           `json:"fqdn,omitempty"`
	Health               *ProtectionHealth `json:"health,omitempty"`
	HealthErrors         *[]HealthError    `json:"healthErrors,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	LastDiscoveryInUtc   *string           `json:"lastDiscoveryInUtc,omitempty"`
	LastHeartbeatUtc     *string           `json:"lastHeartbeatUtc,omitempty"`
	Name                 *string           `json:"name,omitempty"`
	ProtectedItemCount   *int64            `json:"protectedItemCount,omitempty"`
	VcenterId            *string           `json:"vcenterId,omitempty"`
	Version              *string           `json:"version,omitempty"`
}
