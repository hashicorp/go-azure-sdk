package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentResourceRole struct {
	AccessPackageAssignments   *[]AccessPackageAssignment  `json:"accessPackageAssignments,omitempty"`
	AccessPackageResourceRole  *AccessPackageResourceRole  `json:"accessPackageResourceRole,omitempty"`
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
	AccessPackageSubject       *AccessPackageSubject       `json:"accessPackageSubject,omitempty"`
	Id                         *string                     `json:"id,omitempty"`
	ODataType                  *string                     `json:"@odata.type,omitempty"`
	OriginId                   *string                     `json:"originId,omitempty"`
	OriginSystem               *string                     `json:"originSystem,omitempty"`
	Status                     *string                     `json:"status,omitempty"`
}
