package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalGroups struct {
	AzureLogAnalytics     *LogAnalyticsSignals          `json:"azureLogAnalytics,omitempty"`
	AzureMonitorWorkspace *AzureMonitorWorkspaceSignals `json:"azureMonitorWorkspace,omitempty"`
	AzureResource         *AzureResourceSignals         `json:"azureResource,omitempty"`
	Dependencies          *DependenciesSignalGroupV2    `json:"dependencies,omitempty"`
	External              *ExternalSignalGroup          `json:"external,omitempty"`
}
