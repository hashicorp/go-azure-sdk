package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedIdentity struct {
	Details      *DetailsInfo `json:"details,omitempty"`
	DisplayName  *string      `json:"displayName,omitempty"`
	Id           *string      `json:"id,omitempty"`
	IdentityType *string      `json:"identityType,omitempty"`
	ODataType    *string      `json:"@odata.type,omitempty"`
}
