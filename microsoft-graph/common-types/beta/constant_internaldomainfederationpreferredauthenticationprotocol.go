package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationPreferredAuthenticationProtocol string

const (
	InternalDomainFederationPreferredAuthenticationProtocol_Saml  InternalDomainFederationPreferredAuthenticationProtocol = "saml"
	InternalDomainFederationPreferredAuthenticationProtocol_WsFed InternalDomainFederationPreferredAuthenticationProtocol = "wsFed"
)

func PossibleValuesForInternalDomainFederationPreferredAuthenticationProtocol() []string {
	return []string{
		string(InternalDomainFederationPreferredAuthenticationProtocol_Saml),
		string(InternalDomainFederationPreferredAuthenticationProtocol_WsFed),
	}
}

func (s *InternalDomainFederationPreferredAuthenticationProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInternalDomainFederationPreferredAuthenticationProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInternalDomainFederationPreferredAuthenticationProtocol(input string) (*InternalDomainFederationPreferredAuthenticationProtocol, error) {
	vals := map[string]InternalDomainFederationPreferredAuthenticationProtocol{
		"saml":  InternalDomainFederationPreferredAuthenticationProtocol_Saml,
		"wsfed": InternalDomainFederationPreferredAuthenticationProtocol_WsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationPreferredAuthenticationProtocol(input)
	return &out, nil
}
