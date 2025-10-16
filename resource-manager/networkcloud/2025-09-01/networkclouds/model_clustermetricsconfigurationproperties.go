package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterMetricsConfigurationProperties struct {
	CollectionInterval    int64                                         `json:"collectionInterval"`
	DetailedStatus        *ClusterMetricsConfigurationDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                                       `json:"detailedStatusMessage,omitempty"`
	DisabledMetrics       *[]string                                     `json:"disabledMetrics,omitempty"`
	EnabledMetrics        *[]string                                     `json:"enabledMetrics,omitempty"`
	ProvisioningState     *ClusterMetricsConfigurationProvisioningState `json:"provisioningState,omitempty"`
}
