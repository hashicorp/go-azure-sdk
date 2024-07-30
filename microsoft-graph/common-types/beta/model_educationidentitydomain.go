package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityDomain struct {
	AppliesTo *EducationIdentityDomainAppliesTo `json:"appliesTo,omitempty"`
	Name      *string                           `json:"name,omitempty"`
	ODataType *string                           `json:"@odata.type,omitempty"`
}
