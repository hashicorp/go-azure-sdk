package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterface struct {
	Description      *string `json:"description,omitempty"`
	IpV4Address      *string `json:"ipV4Address,omitempty"`
	IpV6Address      *string `json:"ipV6Address,omitempty"`
	LocalIpV6Address *string `json:"localIpV6Address,omitempty"`
	MacAddress       *string `json:"macAddress,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
