package networkclouds

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/zones"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolProperties struct {
	AdministratorConfiguration   *AdministratorConfiguration   `json:"administratorConfiguration,omitempty"`
	AgentOptions                 *AgentOptions                 `json:"agentOptions,omitempty"`
	AttachedNetworkConfiguration *AttachedNetworkConfiguration `json:"attachedNetworkConfiguration,omitempty"`
	AvailabilityZones            *zones.Schema                 `json:"availabilityZones,omitempty"`
	Count                        int64                         `json:"count"`
	DetailedStatus               *AgentPoolDetailedStatus      `json:"detailedStatus,omitempty"`
	DetailedStatusMessage        *string                       `json:"detailedStatusMessage,omitempty"`
	KubernetesVersion            *string                       `json:"kubernetesVersion,omitempty"`
	Labels                       *[]KubernetesLabel            `json:"labels,omitempty"`
	Mode                         AgentPoolMode                 `json:"mode"`
	ProvisioningState            *AgentPoolProvisioningState   `json:"provisioningState,omitempty"`
	Taints                       *[]KubernetesLabel            `json:"taints,omitempty"`
	UpgradeSettings              *AgentPoolUpgradeSettings     `json:"upgradeSettings,omitempty"`
	VMSkuName                    string                        `json:"vmSkuName"`
}
