package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateNetworkSiblingSetRequest struct {
	NetworkFeatures          NetworkFeatures `json:"networkFeatures"`
	NetworkSiblingSetId      string          `json:"networkSiblingSetId"`
	NetworkSiblingSetStateId string          `json:"networkSiblingSetStateId"`
	SubnetId                 string          `json:"subnetId"`
}
