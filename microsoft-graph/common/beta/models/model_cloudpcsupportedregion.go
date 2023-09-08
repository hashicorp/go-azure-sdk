package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegion struct {
	DisplayName       *string                                  `json:"displayName,omitempty"`
	Id                *string                                  `json:"id,omitempty"`
	ODataType         *string                                  `json:"@odata.type,omitempty"`
	RegionGroup       *CloudPCSupportedRegionRegionGroup       `json:"regionGroup,omitempty"`
	RegionStatus      *CloudPCSupportedRegionRegionStatus      `json:"regionStatus,omitempty"`
	SupportedSolution *CloudPCSupportedRegionSupportedSolution `json:"supportedSolution,omitempty"`
}
