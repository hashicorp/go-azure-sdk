package changedatacaptureconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureConfigurationProperties struct {
	DatabaseName          string                                 `json:"databaseName"`
	DesiredState          ChangeDataCaptureDesiredState          `json:"desiredState"`
	DestinationResourceId string                                 `json:"destinationResourceId"`
	EffectiveState        ChangeDataCaptureEffectiveState        `json:"effectiveState"`
	Extension             ChangeDataCaptureExtensionStatus       `json:"extension"`
	LogicalStreamId       string                                 `json:"logicalStreamId"`
	ProvisioningState     *ProvisioningState                     `json:"provisioningState,omitempty"`
	Replication           ChangeDataCaptureReplicationStatus     `json:"replication"`
	Status                ChangeDataCaptureStatus                `json:"status"`
	Topology              ChangeDataCaptureTopologyBindingStatus `json:"topology"`
	UsageProfile          ChangeDataCaptureUsageProfile          `json:"usageProfile"`
}
