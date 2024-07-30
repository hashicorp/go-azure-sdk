package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceScope struct {
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`
	Description           *string                `json:"description,omitempty"`
	DisplayName           *string                `json:"displayName,omitempty"`
	Id                    *string                `json:"id,omitempty"`
	IsRootScope           *bool                  `json:"isRootScope,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	OriginId              *string                `json:"originId,omitempty"`
	OriginSystem          *string                `json:"originSystem,omitempty"`
	RoleOriginId          *string                `json:"roleOriginId,omitempty"`
	Url                   *string                `json:"url,omitempty"`
}
