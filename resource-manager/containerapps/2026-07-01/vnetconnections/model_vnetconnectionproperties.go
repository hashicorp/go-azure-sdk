package vnetconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetConnectionProperties struct {
	ProvisioningState *VnetConnectionProvisioningState `json:"provisioningState,omitempty"`
	SubnetId          *string                          `json:"subnetId,omitempty"`
}
