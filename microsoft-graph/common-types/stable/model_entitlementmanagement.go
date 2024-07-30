package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagement struct {
	AccessPackageAssignmentApprovals *[]Approval                         `json:"accessPackageAssignmentApprovals,omitempty"`
	AccessPackages                   *[]AccessPackage                    `json:"accessPackages,omitempty"`
	AssignmentPolicies               *[]AccessPackageAssignmentPolicy    `json:"assignmentPolicies,omitempty"`
	AssignmentRequests               *[]AccessPackageAssignmentRequest   `json:"assignmentRequests,omitempty"`
	Assignments                      *[]AccessPackageAssignment          `json:"assignments,omitempty"`
	Catalogs                         *[]AccessPackageCatalog             `json:"catalogs,omitempty"`
	ConnectedOrganizations           *[]ConnectedOrganization            `json:"connectedOrganizations,omitempty"`
	Id                               *string                             `json:"id,omitempty"`
	ODataType                        *string                             `json:"@odata.type,omitempty"`
	ResourceEnvironments             *[]AccessPackageResourceEnvironment `json:"resourceEnvironments,omitempty"`
	ResourceRequests                 *[]AccessPackageResourceRequest     `json:"resourceRequests,omitempty"`
	ResourceRoleScopes               *[]AccessPackageResourceRoleScope   `json:"resourceRoleScopes,omitempty"`
	Resources                        *[]AccessPackageResource            `json:"resources,omitempty"`
	Settings                         *EntitlementManagementSettings      `json:"settings,omitempty"`
}
