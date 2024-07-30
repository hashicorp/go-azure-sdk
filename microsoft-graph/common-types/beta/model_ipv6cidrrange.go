package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPv6CidrRange struct {
	CidrAddress *string `json:"cidrAddress,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
