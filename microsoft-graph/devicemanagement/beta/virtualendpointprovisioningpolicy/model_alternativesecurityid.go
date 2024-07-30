package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlternativeSecurityId struct {
	IdentityProvider *string `json:"identityProvider,omitempty"`
	Key              *string `json:"key,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Type             *int64  `json:"type,omitempty"`
}
