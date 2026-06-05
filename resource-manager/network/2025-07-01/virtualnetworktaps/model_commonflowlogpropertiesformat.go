package virtualnetworktaps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonFlowLogPropertiesFormat struct {
	Enabled                    *bool                             `json:"enabled,omitempty"`
	EnabledFilteringCriteria   *string                           `json:"enabledFilteringCriteria,omitempty"`
	FlowAnalyticsConfiguration *CommonTrafficAnalyticsProperties `json:"flowAnalyticsConfiguration,omitempty"`
	Format                     *CommonFlowLogFormatParameters    `json:"format,omitempty"`
	ProvisioningState          *ProvisioningState                `json:"provisioningState,omitempty"`
	RecordTypes                *string                           `json:"recordTypes,omitempty"`
	RetentionPolicy            *CommonRetentionPolicyParameters  `json:"retentionPolicy,omitempty"`
	StorageId                  string                            `json:"storageId"`
	TargetResourceGuid         *string                           `json:"targetResourceGuid,omitempty"`
	TargetResourceId           string                            `json:"targetResourceId"`
}
