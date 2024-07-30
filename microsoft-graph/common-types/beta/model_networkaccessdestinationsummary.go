package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestinationSummary struct {
	Count       *int64                                      `json:"count,omitempty"`
	Destination *string                                     `json:"destination,omitempty"`
	ODataType   *string                                     `json:"@odata.type,omitempty"`
	TrafficType *NetworkaccessDestinationSummaryTrafficType `json:"trafficType,omitempty"`
}
