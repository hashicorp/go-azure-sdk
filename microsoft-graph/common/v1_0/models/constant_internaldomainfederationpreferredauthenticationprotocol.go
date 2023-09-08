package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationPreferredAuthenticationProtocol string

const (
	InternalDomainFederationPreferredAuthenticationProtocolsaml  InternalDomainFederationPreferredAuthenticationProtocol = "Saml"
	InternalDomainFederationPreferredAuthenticationProtocolwsFed InternalDomainFederationPreferredAuthenticationProtocol = "WsFed"
)

func PossibleValuesForInternalDomainFederationPreferredAuthenticationProtocol() []string {
	return []string{
		string(InternalDomainFederationPreferredAuthenticationProtocolsaml),
		string(InternalDomainFederationPreferredAuthenticationProtocolwsFed),
	}
}

func parseInternalDomainFederationPreferredAuthenticationProtocol(input string) (*InternalDomainFederationPreferredAuthenticationProtocol, error) {
	vals := map[string]InternalDomainFederationPreferredAuthenticationProtocol{
		"saml":  InternalDomainFederationPreferredAuthenticationProtocolsaml,
		"wsfed": InternalDomainFederationPreferredAuthenticationProtocolwsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationPreferredAuthenticationProtocol(input)
	return &out, nil
}
