package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRoleScope struct {
	AccessPackageResourceRole  *AccessPackageResourceRole  `json:"accessPackageResourceRole,omitempty"`
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
	CreatedBy                  *string                     `json:"createdBy,omitempty"`
	CreatedDateTime            *string                     `json:"createdDateTime,omitempty"`
	Id                         *string                     `json:"id,omitempty"`
	ModifiedBy                 *string                     `json:"modifiedBy,omitempty"`
	ModifiedDateTime           *string                     `json:"modifiedDateTime,omitempty"`
	ODataType                  *string                     `json:"@odata.type,omitempty"`
}
