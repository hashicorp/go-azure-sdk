package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevOpsConfigurationProperties struct {
	Authorization                   *Authorization           `json:"authorization,omitempty"`
	AutoDiscovery                   *AutoDiscovery           `json:"autoDiscovery,omitempty"`
	Capabilities                    *[]DevOpsCapability      `json:"capabilities,omitempty"`
	ProvisioningState               *DevOpsProvisioningState `json:"provisioningState,omitempty"`
	ProvisioningStatusMessage       *string                  `json:"provisioningStatusMessage,omitempty"`
	ProvisioningStatusUpdateTimeUtc *string                  `json:"provisioningStatusUpdateTimeUtc,omitempty"`
	TopLevelInventoryList           *[]string                `json:"topLevelInventoryList,omitempty"`
}
