package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationFederatedIdpMfaBehavior string

const (
	InternalDomainFederationFederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp InternalDomainFederationFederatedIdpMfaBehavior = "acceptIfMfaDoneByFederatedIdp"
	InternalDomainFederationFederatedIdpMfaBehavior_EnforceMfaByFederatedIdp      InternalDomainFederationFederatedIdpMfaBehavior = "enforceMfaByFederatedIdp"
	InternalDomainFederationFederatedIdpMfaBehavior_RejectMfaByFederatedIdp       InternalDomainFederationFederatedIdpMfaBehavior = "rejectMfaByFederatedIdp"
)

func PossibleValuesForInternalDomainFederationFederatedIdpMfaBehavior() []string {
	return []string{
		string(InternalDomainFederationFederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp),
		string(InternalDomainFederationFederatedIdpMfaBehavior_EnforceMfaByFederatedIdp),
		string(InternalDomainFederationFederatedIdpMfaBehavior_RejectMfaByFederatedIdp),
	}
}

func (s *InternalDomainFederationFederatedIdpMfaBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInternalDomainFederationFederatedIdpMfaBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInternalDomainFederationFederatedIdpMfaBehavior(input string) (*InternalDomainFederationFederatedIdpMfaBehavior, error) {
	vals := map[string]InternalDomainFederationFederatedIdpMfaBehavior{
		"acceptifmfadonebyfederatedidp": InternalDomainFederationFederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp,
		"enforcemfabyfederatedidp":      InternalDomainFederationFederatedIdpMfaBehavior_EnforceMfaByFederatedIdp,
		"rejectmfabyfederatedidp":       InternalDomainFederationFederatedIdpMfaBehavior_RejectMfaByFederatedIdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationFederatedIdpMfaBehavior(input)
	return &out, nil
}
