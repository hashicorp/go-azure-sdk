package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDomainConfiguration struct {
	CertificatePassword        *string `json:"certificatePassword,omitempty"`
	CertificateValue           *string `json:"certificateValue,omitempty"`
	CustomDomainVerificationId *string `json:"customDomainVerificationId,omitempty"`
	DnsSuffix                  *string `json:"dnsSuffix,omitempty"`
	ExpirationDate             *string `json:"expirationDate,omitempty"`
	SubjectName                *string `json:"subjectName,omitempty"`
	Thumbprint                 *string `json:"thumbprint,omitempty"`
}
