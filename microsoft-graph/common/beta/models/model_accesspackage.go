package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackage struct {
	AccessPackageAssignmentPolicies *[]AccessPackageAssignmentPolicy  `json:"accessPackageAssignmentPolicies,omitempty"`
	AccessPackageCatalog            *AccessPackageCatalog             `json:"accessPackageCatalog,omitempty"`
	AccessPackageResourceRoleScopes *[]AccessPackageResourceRoleScope `json:"accessPackageResourceRoleScopes,omitempty"`
	AccessPackagesIncompatibleWith  *[]AccessPackage                  `json:"accessPackagesIncompatibleWith,omitempty"`
	CatalogId                       *string                           `json:"catalogId,omitempty"`
	CreatedBy                       *string                           `json:"createdBy,omitempty"`
	CreatedDateTime                 *string                           `json:"createdDateTime,omitempty"`
	Description                     *string                           `json:"description,omitempty"`
	DisplayName                     *string                           `json:"displayName,omitempty"`
	Id                              *string                           `json:"id,omitempty"`
	IncompatibleAccessPackages      *[]AccessPackage                  `json:"incompatibleAccessPackages,omitempty"`
	IncompatibleGroups              *[]Group                          `json:"incompatibleGroups,omitempty"`
	IsHidden                        *bool                             `json:"isHidden,omitempty"`
	IsRoleScopesVisible             *bool                             `json:"isRoleScopesVisible,omitempty"`
	ModifiedBy                      *string                           `json:"modifiedBy,omitempty"`
	ModifiedDateTime                *string                           `json:"modifiedDateTime,omitempty"`
	ODataType                       *string                           `json:"@odata.type,omitempty"`
}
