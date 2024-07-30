package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalog struct {
	AccessPackageCustomWorkflowExtensions *[]CustomCalloutExtension               `json:"accessPackageCustomWorkflowExtensions,omitempty"`
	AccessPackageResourceRoles            *[]AccessPackageResourceRole            `json:"accessPackageResourceRoles,omitempty"`
	AccessPackageResourceScopes           *[]AccessPackageResourceScope           `json:"accessPackageResourceScopes,omitempty"`
	AccessPackageResources                *[]AccessPackageResource                `json:"accessPackageResources,omitempty"`
	AccessPackages                        *[]AccessPackage                        `json:"accessPackages,omitempty"`
	CatalogStatus                         *string                                 `json:"catalogStatus,omitempty"`
	CatalogType                           *string                                 `json:"catalogType,omitempty"`
	CreatedBy                             *string                                 `json:"createdBy,omitempty"`
	CreatedDateTime                       *string                                 `json:"createdDateTime,omitempty"`
	CustomAccessPackageWorkflowExtensions *[]CustomAccessPackageWorkflowExtension `json:"customAccessPackageWorkflowExtensions,omitempty"`
	Description                           *string                                 `json:"description,omitempty"`
	DisplayName                           *string                                 `json:"displayName,omitempty"`
	Id                                    *string                                 `json:"id,omitempty"`
	IsExternallyVisible                   *bool                                   `json:"isExternallyVisible,omitempty"`
	ModifiedBy                            *string                                 `json:"modifiedBy,omitempty"`
	ModifiedDateTime                      *string                                 `json:"modifiedDateTime,omitempty"`
	ODataType                             *string                                 `json:"@odata.type,omitempty"`
}
