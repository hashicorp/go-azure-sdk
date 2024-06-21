package integrationruntime

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfHostedIntegrationRuntimeNode struct {
	Capabilities        *map[string]string                      `json:"capabilities,omitempty"`
	ConcurrentJobsLimit *int64                                  `json:"concurrentJobsLimit,omitempty"`
	ExpiryTime          *string                                 `json:"expiryTime,omitempty"`
	HostServiceUri      *string                                 `json:"hostServiceUri,omitempty"`
	IsActiveDispatcher  *bool                                   `json:"isActiveDispatcher,omitempty"`
	LastConnectTime     *string                                 `json:"lastConnectTime,omitempty"`
	LastEndUpdateTime   *string                                 `json:"lastEndUpdateTime,omitempty"`
	LastStartTime       *string                                 `json:"lastStartTime,omitempty"`
	LastStartUpdateTime *string                                 `json:"lastStartUpdateTime,omitempty"`
	LastStopTime        *string                                 `json:"lastStopTime,omitempty"`
	LastUpdateResult    *IntegrationRuntimeUpdateResult         `json:"lastUpdateResult,omitempty"`
	MachineName         *string                                 `json:"machineName,omitempty"`
	MaxConcurrentJobs   *int64                                  `json:"maxConcurrentJobs,omitempty"`
	NodeName            *string                                 `json:"nodeName,omitempty"`
	RegisterTime        *string                                 `json:"registerTime,omitempty"`
	Status              *SelfHostedIntegrationRuntimeNodeStatus `json:"status,omitempty"`
	Version             *string                                 `json:"version,omitempty"`
	VersionStatus       *string                                 `json:"versionStatus,omitempty"`
}
