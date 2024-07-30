package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProxiedDomain struct {
	IpAddressOrFQDN *string `json:"ipAddressOrFQDN,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	Proxy           *string `json:"proxy,omitempty"`
}
