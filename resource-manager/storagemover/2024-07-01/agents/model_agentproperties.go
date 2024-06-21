package agents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentProperties struct {
	AgentStatus         *AgentStatus                 `json:"agentStatus,omitempty"`
	AgentVersion        *string                      `json:"agentVersion,omitempty"`
	ArcResourceId       string                       `json:"arcResourceId"`
	ArcVMUuid           string                       `json:"arcVmUuid"`
	Description         *string                      `json:"description,omitempty"`
	ErrorDetails        *AgentPropertiesErrorDetails `json:"errorDetails,omitempty"`
	LastStatusUpdate    *string                      `json:"lastStatusUpdate,omitempty"`
	LocalIPAddress      *string                      `json:"localIPAddress,omitempty"`
	MemoryInMB          *int64                       `json:"memoryInMB,omitempty"`
	NumberOfCores       *int64                       `json:"numberOfCores,omitempty"`
	ProvisioningState   *ProvisioningState           `json:"provisioningState,omitempty"`
	TimeZone            *string                      `json:"timeZone,omitempty"`
	UploadLimitSchedule *UploadLimitSchedule         `json:"uploadLimitSchedule,omitempty"`
	UptimeInSeconds     *int64                       `json:"uptimeInSeconds,omitempty"`
}
