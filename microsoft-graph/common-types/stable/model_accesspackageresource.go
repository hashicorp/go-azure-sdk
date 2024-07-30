package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResource struct {
	Attributes       *[]AccessPackageResourceAttribute `json:"attributes,omitempty"`
	CreatedDateTime  *string                           `json:"createdDateTime,omitempty"`
	Description      *string                           `json:"description,omitempty"`
	DisplayName      *string                           `json:"displayName,omitempty"`
	Environment      *AccessPackageResourceEnvironment `json:"environment,omitempty"`
	Id               *string                           `json:"id,omitempty"`
	ModifiedDateTime *string                           `json:"modifiedDateTime,omitempty"`
	ODataType        *string                           `json:"@odata.type,omitempty"`
	OriginId         *string                           `json:"originId,omitempty"`
	OriginSystem     *string                           `json:"originSystem,omitempty"`
	Roles            *[]AccessPackageResourceRole      `json:"roles,omitempty"`
	Scopes           *[]AccessPackageResourceScope     `json:"scopes,omitempty"`
}
