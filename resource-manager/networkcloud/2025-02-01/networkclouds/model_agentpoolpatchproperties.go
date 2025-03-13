package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolPatchProperties struct {
	AdministratorConfiguration *NodePoolAdministratorConfigurationPatch `json:"administratorConfiguration,omitempty"`
	Count                      *int64                                   `json:"count,omitempty"`
	UpgradeSettings            *AgentPoolUpgradeSettings                `json:"upgradeSettings,omitempty"`
}
