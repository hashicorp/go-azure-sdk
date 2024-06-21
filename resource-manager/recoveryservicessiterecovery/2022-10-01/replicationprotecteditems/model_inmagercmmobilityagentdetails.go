package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InMageRcmMobilityAgentDetails struct {
	AgentVersionExpiryDate               *string                      `json:"agentVersionExpiryDate,omitempty"`
	DriverVersion                        *string                      `json:"driverVersion,omitempty"`
	DriverVersionExpiryDate              *string                      `json:"driverVersionExpiryDate,omitempty"`
	IsUpgradeable                        *string                      `json:"isUpgradeable,omitempty"`
	LastHeartbeatUtc                     *string                      `json:"lastHeartbeatUtc,omitempty"`
	LatestAgentReleaseDate               *string                      `json:"latestAgentReleaseDate,omitempty"`
	LatestUpgradableVersionWithoutReboot *string                      `json:"latestUpgradableVersionWithoutReboot,omitempty"`
	LatestVersion                        *string                      `json:"latestVersion,omitempty"`
	ReasonsBlockingUpgrade               *[]AgentUpgradeBlockedReason `json:"reasonsBlockingUpgrade,omitempty"`
	Version                              *string                      `json:"version,omitempty"`
}
