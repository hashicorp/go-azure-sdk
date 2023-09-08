package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol string

const (
	SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolsaml  SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol = "Saml"
	SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolwsFed SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol = "WsFed"
)

func PossibleValuesForSamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol() []string {
	return []string{
		string(SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolsaml),
		string(SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolwsFed),
	}
}

func parseSamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol(input string) (*SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol, error) {
	vals := map[string]SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol{
		"saml":  SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolsaml,
		"wsfed": SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocolwsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol(input)
	return &out, nil
}
