package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRoleScope struct {
	CreatedDateTime *string                     `json:"createdDateTime,omitempty"`
	Id              *string                     `json:"id,omitempty"`
	ODataType       *string                     `json:"@odata.type,omitempty"`
	Role            *AccessPackageResourceRole  `json:"role,omitempty"`
	Scope           *AccessPackageResourceScope `json:"scope,omitempty"`
}
