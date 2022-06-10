package autoscalevcores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScaleVCoreProperties struct {
	CapacityLimit     *int64                  `json:"capacityLimit,omitempty"`
	CapacityObjectId  *string                 `json:"capacityObjectId,omitempty"`
	ProvisioningState *VCoreProvisioningState `json:"provisioningState,omitempty"`
}
