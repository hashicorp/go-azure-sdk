package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalog struct {
	AccessPackages           *[]AccessPackage                 `json:"accessPackages,omitempty"`
	CatalogType              *AccessPackageCatalogCatalogType `json:"catalogType,omitempty"`
	CreatedDateTime          *string                          `json:"createdDateTime,omitempty"`
	CustomWorkflowExtensions *[]CustomCalloutExtension        `json:"customWorkflowExtensions,omitempty"`
	Description              *string                          `json:"description,omitempty"`
	DisplayName              *string                          `json:"displayName,omitempty"`
	Id                       *string                          `json:"id,omitempty"`
	IsExternallyVisible      *bool                            `json:"isExternallyVisible,omitempty"`
	ModifiedDateTime         *string                          `json:"modifiedDateTime,omitempty"`
	ODataType                *string                          `json:"@odata.type,omitempty"`
	ResourceRoles            *[]AccessPackageResourceRole     `json:"resourceRoles,omitempty"`
	ResourceScopes           *[]AccessPackageResourceScope    `json:"resourceScopes,omitempty"`
	Resources                *[]AccessPackageResource         `json:"resources,omitempty"`
	State                    *AccessPackageCatalogState       `json:"state,omitempty"`
}
