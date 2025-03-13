package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentlessConfiguration struct {
	AgentlessAutoDiscovery *AutoDiscovery       `json:"agentlessAutoDiscovery,omitempty"`
	AgentlessEnabled       *AgentlessEnablement `json:"agentlessEnabled,omitempty"`
	InventoryList          *[]InventoryList     `json:"inventoryList,omitempty"`
	InventoryListType      *InventoryListKind   `json:"inventoryListType,omitempty"`
	Scanners               *[]string            `json:"scanners,omitempty"`
}
