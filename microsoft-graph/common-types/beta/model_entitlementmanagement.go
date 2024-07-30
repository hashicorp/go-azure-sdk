package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagement struct {
	AccessPackageAssignmentApprovals     *[]Approval                            `json:"accessPackageAssignmentApprovals,omitempty"`
	AccessPackageAssignmentPolicies      *[]AccessPackageAssignmentPolicy       `json:"accessPackageAssignmentPolicies,omitempty"`
	AccessPackageAssignmentRequests      *[]AccessPackageAssignmentRequest      `json:"accessPackageAssignmentRequests,omitempty"`
	AccessPackageAssignmentResourceRoles *[]AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`
	AccessPackageAssignments             *[]AccessPackageAssignment             `json:"accessPackageAssignments,omitempty"`
	AccessPackageCatalogs                *[]AccessPackageCatalog                `json:"accessPackageCatalogs,omitempty"`
	AccessPackageResourceEnvironments    *[]AccessPackageResourceEnvironment    `json:"accessPackageResourceEnvironments,omitempty"`
	AccessPackageResourceRequests        *[]AccessPackageResourceRequest        `json:"accessPackageResourceRequests,omitempty"`
	AccessPackageResourceRoleScopes      *[]AccessPackageResourceRoleScope      `json:"accessPackageResourceRoleScopes,omitempty"`
	AccessPackageResources               *[]AccessPackageResource               `json:"accessPackageResources,omitempty"`
	AccessPackages                       *[]AccessPackage                       `json:"accessPackages,omitempty"`
	ConnectedOrganizations               *[]ConnectedOrganization               `json:"connectedOrganizations,omitempty"`
	Id                                   *string                                `json:"id,omitempty"`
	ODataType                            *string                                `json:"@odata.type,omitempty"`
	Settings                             *EntitlementManagementSettings         `json:"settings,omitempty"`
	Subjects                             *[]AccessPackageSubject                `json:"subjects,omitempty"`
}
