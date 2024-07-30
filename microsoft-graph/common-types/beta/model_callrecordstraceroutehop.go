package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsTraceRouteHop struct {
	HopCount      *int64  `json:"hopCount,omitempty"`
	IpAddress     *string `json:"ipAddress,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	RoundTripTime *string `json:"roundTripTime,omitempty"`
}
