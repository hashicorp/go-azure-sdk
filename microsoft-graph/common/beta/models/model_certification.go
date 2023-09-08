package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Certification struct {
	CertificationDetailsUrl         *string `json:"certificationDetailsUrl,omitempty"`
	CertificationExpirationDateTime *string `json:"certificationExpirationDateTime,omitempty"`
	IsCertifiedByMicrosoft          *bool   `json:"isCertifiedByMicrosoft,omitempty"`
	IsPublisherAttested             *bool   `json:"isPublisherAttested,omitempty"`
	LastCertificationDateTime       *string `json:"lastCertificationDateTime,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
}
