package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RackProperties struct {
	AvailabilityZone      string                 `json:"availabilityZone"`
	ClusterId             *string                `json:"clusterId,omitempty"`
	DetailedStatus        *RackDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                `json:"detailedStatusMessage,omitempty"`
	ProvisioningState     *RackProvisioningState `json:"provisioningState,omitempty"`
	RackLocation          string                 `json:"rackLocation"`
	RackSerialNumber      string                 `json:"rackSerialNumber"`
	RackSkuId             string                 `json:"rackSkuId"`
}
