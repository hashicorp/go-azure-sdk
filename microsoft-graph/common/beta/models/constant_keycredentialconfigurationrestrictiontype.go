package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredentialConfigurationRestrictionType string

const (
	KeyCredentialConfigurationRestrictionTypeasymmetricKeyLifetime       KeyCredentialConfigurationRestrictionType = "AsymmetricKeyLifetime"
	KeyCredentialConfigurationRestrictionTypetrustedCertificateAuthority KeyCredentialConfigurationRestrictionType = "TrustedCertificateAuthority"
)

func PossibleValuesForKeyCredentialConfigurationRestrictionType() []string {
	return []string{
		string(KeyCredentialConfigurationRestrictionTypeasymmetricKeyLifetime),
		string(KeyCredentialConfigurationRestrictionTypetrustedCertificateAuthority),
	}
}

func parseKeyCredentialConfigurationRestrictionType(input string) (*KeyCredentialConfigurationRestrictionType, error) {
	vals := map[string]KeyCredentialConfigurationRestrictionType{
		"asymmetrickeylifetime":       KeyCredentialConfigurationRestrictionTypeasymmetricKeyLifetime,
		"trustedcertificateauthority": KeyCredentialConfigurationRestrictionTypetrustedCertificateAuthority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyCredentialConfigurationRestrictionType(input)
	return &out, nil
}
