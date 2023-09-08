package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResource struct {
	AccessPackageResourceEnvironment *AccessPackageResourceEnvironment `json:"accessPackageResourceEnvironment,omitempty"`
	AccessPackageResourceRoles       *[]AccessPackageResourceRole      `json:"accessPackageResourceRoles,omitempty"`
	AccessPackageResourceScopes      *[]AccessPackageResourceScope     `json:"accessPackageResourceScopes,omitempty"`
	AddedBy                          *string                           `json:"addedBy,omitempty"`
	AddedOn                          *string                           `json:"addedOn,omitempty"`
	Attributes                       *[]AccessPackageResourceAttribute `json:"attributes,omitempty"`
	Description                      *string                           `json:"description,omitempty"`
	DisplayName                      *string                           `json:"displayName,omitempty"`
	Id                               *string                           `json:"id,omitempty"`
	IsPendingOnboarding              *bool                             `json:"isPendingOnboarding,omitempty"`
	ODataType                        *string                           `json:"@odata.type,omitempty"`
	OriginId                         *string                           `json:"originId,omitempty"`
	OriginSystem                     *string                           `json:"originSystem,omitempty"`
	ResourceType                     *string                           `json:"resourceType,omitempty"`
	Url                              *string                           `json:"url,omitempty"`
}
