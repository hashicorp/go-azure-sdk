package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationLicenseAssignment struct {
	AppliesTo *EducationSynchronizationLicenseAssignmentAppliesTo `json:"appliesTo,omitempty"`
	ODataType *string                                             `json:"@odata.type,omitempty"`
	SkuIds    *[]string                                           `json:"skuIds,omitempty"`
}
