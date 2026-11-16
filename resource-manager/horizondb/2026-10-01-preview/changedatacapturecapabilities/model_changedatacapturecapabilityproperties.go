package changedatacapturecapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureCapabilityProperties struct {
	AutoUpgradeExtension    *bool                                    `json:"autoUpgradeExtension,omitempty"`
	ConfigurationCount      int64                                    `json:"configurationCount"`
	DesiredState            ChangeDataCaptureDesiredState            `json:"desiredState"`
	DestinationCount        int64                                    `json:"destinationCount"`
	EffectiveState          ChangeDataCaptureEffectiveState          `json:"effectiveState"`
	Extension               ChangeDataCaptureExtensionStatus         `json:"extension"`
	Prerequisites           ChangeDataCapturePrerequisiteStatus      `json:"prerequisites"`
	ProvisioningState       *ProvisioningState                       `json:"provisioningState,omitempty"`
	ReplicationArchitecture ChangeDataCaptureReplicationArchitecture `json:"replicationArchitecture"`
	Status                  ChangeDataCaptureStatus                  `json:"status"`
}
