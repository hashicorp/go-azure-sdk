package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrincipalResourceMembershipsScope struct {
	ODataType       *string              `json:"@odata.type,omitempty"`
	PrincipalScopes *[]AccessReviewScope `json:"principalScopes,omitempty"`
	ResourceScopes  *[]AccessReviewScope `json:"resourceScopes,omitempty"`
}
