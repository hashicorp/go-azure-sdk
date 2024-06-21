package replicationfabrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarsAgentDetails struct {
	BiosId           *string           `json:"biosId,omitempty"`
	FabricObjectId   *string           `json:"fabricObjectId,omitempty"`
	Fqdn             *string           `json:"fqdn,omitempty"`
	Health           *ProtectionHealth `json:"health,omitempty"`
	HealthErrors     *[]HealthError    `json:"healthErrors,omitempty"`
	Id               *string           `json:"id,omitempty"`
	LastHeartbeatUtc *string           `json:"lastHeartbeatUtc,omitempty"`
	Name             *string           `json:"name,omitempty"`
	Version          *string           `json:"version,omitempty"`
}
