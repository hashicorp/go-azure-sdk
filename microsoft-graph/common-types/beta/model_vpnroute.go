package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnRoute struct {
	DestinationPrefix *string `json:"destinationPrefix,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	PrefixSize        *int64  `json:"prefixSize,omitempty"`
}
