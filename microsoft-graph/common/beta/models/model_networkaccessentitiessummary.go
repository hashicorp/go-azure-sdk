package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEntitiesSummary struct {
	DeviceCount   *int64                                   `json:"deviceCount,omitempty"`
	ODataType     *string                                  `json:"@odata.type,omitempty"`
	TrafficType   *NetworkaccessEntitiesSummaryTrafficType `json:"trafficType,omitempty"`
	UserCount     *int64                                   `json:"userCount,omitempty"`
	WorkloadCount *int64                                   `json:"workloadCount,omitempty"`
}
