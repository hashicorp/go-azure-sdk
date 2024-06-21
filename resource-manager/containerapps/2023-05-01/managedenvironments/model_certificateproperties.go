package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateProperties struct {
	ExpirationDate          *string                       `json:"expirationDate,omitempty"`
	IssueDate               *string                       `json:"issueDate,omitempty"`
	Issuer                  *string                       `json:"issuer,omitempty"`
	Password                *string                       `json:"password,omitempty"`
	ProvisioningState       *CertificateProvisioningState `json:"provisioningState,omitempty"`
	PublicKeyHash           *string                       `json:"publicKeyHash,omitempty"`
	SubjectAlternativeNames *[]string                     `json:"subjectAlternativeNames,omitempty"`
	SubjectName             *string                       `json:"subjectName,omitempty"`
	Thumbprint              *string                       `json:"thumbprint,omitempty"`
	Valid                   *bool                         `json:"valid,omitempty"`
	Value                   *string                       `json:"value,omitempty"`
}
