package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterNode struct {
	AgentPoolId           *string                              `json:"agentPoolId,omitempty"`
	AvailabilityZone      *string                              `json:"availabilityZone,omitempty"`
	BareMetalMachineId    *string                              `json:"bareMetalMachineId,omitempty"`
	CpuCores              *int64                               `json:"cpuCores,omitempty"`
	DetailedStatus        *KubernetesClusterNodeDetailedStatus `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                              `json:"detailedStatusMessage,omitempty"`
	DiskSizeGB            *int64                               `json:"diskSizeGB,omitempty"`
	Image                 *string                              `json:"image,omitempty"`
	KubernetesVersion     *string                              `json:"kubernetesVersion,omitempty"`
	Labels                *[]KubernetesLabel                   `json:"labels,omitempty"`
	MemorySizeGB          *int64                               `json:"memorySizeGB,omitempty"`
	Mode                  *AgentPoolMode                       `json:"mode,omitempty"`
	Name                  *string                              `json:"name,omitempty"`
	NetworkAttachments    *[]NetworkAttachment                 `json:"networkAttachments,omitempty"`
	PowerState            *KubernetesNodePowerState            `json:"powerState,omitempty"`
	Role                  *KubernetesNodeRole                  `json:"role,omitempty"`
	Taints                *[]KubernetesLabel                   `json:"taints,omitempty"`
	VMSkuName             *string                              `json:"vmSkuName,omitempty"`
}
