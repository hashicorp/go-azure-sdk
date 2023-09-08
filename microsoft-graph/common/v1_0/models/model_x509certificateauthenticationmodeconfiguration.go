package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationModeConfiguration struct {
	ODataType                                *string                                                                                 `json:"@odata.type,omitempty"`
	Rules                                    *[]X509CertificateRule                                                                  `json:"rules,omitempty"`
	X509CertificateAuthenticationDefaultMode *X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode `json:"x509CertificateAuthenticationDefaultMode,omitempty"`
}
