package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackage struct {
	AccessPackagesIncompatibleWith *[]AccessPackage                  `json:"accessPackagesIncompatibleWith,omitempty"`
	AssignmentPolicies             *[]AccessPackageAssignmentPolicy  `json:"assignmentPolicies,omitempty"`
	Catalog                        *AccessPackageCatalog             `json:"catalog,omitempty"`
	CreatedDateTime                *string                           `json:"createdDateTime,omitempty"`
	Description                    *string                           `json:"description,omitempty"`
	DisplayName                    *string                           `json:"displayName,omitempty"`
	Id                             *string                           `json:"id,omitempty"`
	IncompatibleAccessPackages     *[]AccessPackage                  `json:"incompatibleAccessPackages,omitempty"`
	IncompatibleGroups             *[]Group                          `json:"incompatibleGroups,omitempty"`
	IsHidden                       *bool                             `json:"isHidden,omitempty"`
	ModifiedDateTime               *string                           `json:"modifiedDateTime,omitempty"`
	ODataType                      *string                           `json:"@odata.type,omitempty"`
	ResourceRoleScopes             *[]AccessPackageResourceRoleScope `json:"resourceRoleScopes,omitempty"`
}
