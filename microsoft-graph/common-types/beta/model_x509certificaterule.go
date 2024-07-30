package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateRule struct {
	Identifier                           *string                                                  `json:"identifier,omitempty"`
	IssuerSubjectIdentifier              *string                                                  `json:"issuerSubjectIdentifier,omitempty"`
	ODataType                            *string                                                  `json:"@odata.type,omitempty"`
	PolicyOidIdentifier                  *string                                                  `json:"policyOidIdentifier,omitempty"`
	X509CertificateAuthenticationMode    *X509CertificateRuleX509CertificateAuthenticationMode    `json:"x509CertificateAuthenticationMode,omitempty"`
	X509CertificateRequiredAffinityLevel *X509CertificateRuleX509CertificateRequiredAffinityLevel `json:"x509CertificateRequiredAffinityLevel,omitempty"`
	X509CertificateRuleType              *X509CertificateRuleX509CertificateRuleType              `json:"x509CertificateRuleType,omitempty"`
}
