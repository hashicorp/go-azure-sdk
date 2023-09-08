package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedProviderPreferredAuthenticationProtocol string

const (
	SamlOrWsFedProviderPreferredAuthenticationProtocolsaml  SamlOrWsFedProviderPreferredAuthenticationProtocol = "Saml"
	SamlOrWsFedProviderPreferredAuthenticationProtocolwsFed SamlOrWsFedProviderPreferredAuthenticationProtocol = "WsFed"
)

func PossibleValuesForSamlOrWsFedProviderPreferredAuthenticationProtocol() []string {
	return []string{
		string(SamlOrWsFedProviderPreferredAuthenticationProtocolsaml),
		string(SamlOrWsFedProviderPreferredAuthenticationProtocolwsFed),
	}
}

func parseSamlOrWsFedProviderPreferredAuthenticationProtocol(input string) (*SamlOrWsFedProviderPreferredAuthenticationProtocol, error) {
	vals := map[string]SamlOrWsFedProviderPreferredAuthenticationProtocol{
		"saml":  SamlOrWsFedProviderPreferredAuthenticationProtocolsaml,
		"wsfed": SamlOrWsFedProviderPreferredAuthenticationProtocolwsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlOrWsFedProviderPreferredAuthenticationProtocol(input)
	return &out, nil
}
