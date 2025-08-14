package serviceresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMigrationServiceStatusResponse struct {
	AgentConfiguration *interface{} `json:"agentConfiguration,omitempty"`
	AgentVersion       *string      `json:"agentVersion,omitempty"`
	Status             *string      `json:"status,omitempty"`
	SupportedTaskTypes *[]string    `json:"supportedTaskTypes,omitempty"`
	VMSize             *string      `json:"vmSize,omitempty"`
}
