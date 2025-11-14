package netapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSiblingSet struct {
	NetworkFeatures          *NetworkFeatures                    `json:"networkFeatures,omitempty"`
	NetworkSiblingSetId      *string                             `json:"networkSiblingSetId,omitempty"`
	NetworkSiblingSetStateId *string                             `json:"networkSiblingSetStateId,omitempty"`
	NicInfoList              *[]NicInfo                          `json:"nicInfoList,omitempty"`
	ProvisioningState        *NetworkSiblingSetProvisioningState `json:"provisioningState,omitempty"`
	SubnetId                 *string                             `json:"subnetId,omitempty"`
}
