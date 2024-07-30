package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol string

const (
	SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_Saml  SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol = "saml"
	SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_WsFed SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol = "wsFed"
)

func PossibleValuesForSamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol() []string {
	return []string{
		string(SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_Saml),
		string(SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_WsFed),
	}
}

func (s *SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol(input string) (*SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol, error) {
	vals := map[string]SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol{
		"saml":  SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_Saml,
		"wsfed": SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol_WsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol(input)
	return &out, nil
}
