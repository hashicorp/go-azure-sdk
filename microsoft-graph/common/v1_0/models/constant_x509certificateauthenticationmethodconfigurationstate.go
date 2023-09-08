package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationMethodConfigurationState string

const (
	X509CertificateAuthenticationMethodConfigurationStatedisabled X509CertificateAuthenticationMethodConfigurationState = "Disabled"
	X509CertificateAuthenticationMethodConfigurationStateenabled  X509CertificateAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForX509CertificateAuthenticationMethodConfigurationState() []string {
	return []string{
		string(X509CertificateAuthenticationMethodConfigurationStatedisabled),
		string(X509CertificateAuthenticationMethodConfigurationStateenabled),
	}
}

func parseX509CertificateAuthenticationMethodConfigurationState(input string) (*X509CertificateAuthenticationMethodConfigurationState, error) {
	vals := map[string]X509CertificateAuthenticationMethodConfigurationState{
		"disabled": X509CertificateAuthenticationMethodConfigurationStatedisabled,
		"enabled":  X509CertificateAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateAuthenticationMethodConfigurationState(input)
	return &out, nil
}
