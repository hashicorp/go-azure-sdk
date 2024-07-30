package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceInformation struct {
	CertificationControls *[]CertificationControl `json:"certificationControls,omitempty"`
	CertificationName     *string                 `json:"certificationName,omitempty"`
	ODataType             *string                 `json:"@odata.type,omitempty"`
}
