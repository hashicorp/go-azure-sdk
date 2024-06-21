package certificates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificatePatchResourceProperties struct {
	CanonicalName             *string                    `json:"canonicalName,omitempty"`
	CerBlob                   *string                    `json:"cerBlob,omitempty"`
	DomainValidationMethod    *string                    `json:"domainValidationMethod,omitempty"`
	ExpirationDate            *string                    `json:"expirationDate,omitempty"`
	FriendlyName              *string                    `json:"friendlyName,omitempty"`
	HostNames                 *[]string                  `json:"hostNames,omitempty"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `json:"hostingEnvironmentProfile,omitempty"`
	IssueDate                 *string                    `json:"issueDate,omitempty"`
	Issuer                    *string                    `json:"issuer,omitempty"`
	KeyVaultId                *string                    `json:"keyVaultId,omitempty"`
	KeyVaultSecretName        *string                    `json:"keyVaultSecretName,omitempty"`
	KeyVaultSecretStatus      *KeyVaultSecretStatus      `json:"keyVaultSecretStatus,omitempty"`
	Password                  *string                    `json:"password,omitempty"`
	PfxBlob                   *string                    `json:"pfxBlob,omitempty"`
	PublicKeyHash             *string                    `json:"publicKeyHash,omitempty"`
	SelfLink                  *string                    `json:"selfLink,omitempty"`
	ServerFarmId              *string                    `json:"serverFarmId,omitempty"`
	SiteName                  *string                    `json:"siteName,omitempty"`
	SubjectName               *string                    `json:"subjectName,omitempty"`
	Thumbprint                *string                    `json:"thumbprint,omitempty"`
	Valid                     *bool                      `json:"valid,omitempty"`
}
