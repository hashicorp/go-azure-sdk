package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode string

const (
	X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateMultiFactor  X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode = "X509CertificateMultiFactor"
	X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateSingleFactor X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode = "X509CertificateSingleFactor"
)

func PossibleValuesForX509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode() []string {
	return []string{
		string(X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateMultiFactor),
		string(X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateSingleFactor),
	}
}

func parseX509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode(input string) (*X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode, error) {
	vals := map[string]X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode{
		"x509certificatemultifactor":  X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateMultiFactor,
		"x509certificatesinglefactor": X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultModex509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAuthenticationModeConfigurationX509CertificateAuthenticationDefaultMode(input)
	return &out, nil
}
