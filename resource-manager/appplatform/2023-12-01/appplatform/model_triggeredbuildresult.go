package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredBuildResult struct {
	Id                   *string                                `json:"id,omitempty"`
	Image                *string                                `json:"image,omitempty"`
	LastTransitionReason *string                                `json:"lastTransitionReason,omitempty"`
	LastTransitionStatus *string                                `json:"lastTransitionStatus,omitempty"`
	LastTransitionTime   *string                                `json:"lastTransitionTime,omitempty"`
	ProvisioningState    *TriggeredBuildResultProvisioningState `json:"provisioningState,omitempty"`
}
