package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedCustomDomainCertificatesMetadata struct {
	ExpiryDate  *string `json:"expiryDate,omitempty"`
	IssueDate   *string `json:"issueDate,omitempty"`
	IssuerName  *string `json:"issuerName,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	SubjectName *string `json:"subjectName,omitempty"`
	Thumbprint  *string `json:"thumbprint,omitempty"`
}
