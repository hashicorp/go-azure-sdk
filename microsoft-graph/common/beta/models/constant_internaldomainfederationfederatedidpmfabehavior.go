package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationFederatedIdpMfaBehavior string

const (
	InternalDomainFederationFederatedIdpMfaBehavioracceptIfMfaDoneByFederatedIdp InternalDomainFederationFederatedIdpMfaBehavior = "AcceptIfMfaDoneByFederatedIdp"
	InternalDomainFederationFederatedIdpMfaBehaviorenforceMfaByFederatedIdp      InternalDomainFederationFederatedIdpMfaBehavior = "EnforceMfaByFederatedIdp"
	InternalDomainFederationFederatedIdpMfaBehaviorrejectMfaByFederatedIdp       InternalDomainFederationFederatedIdpMfaBehavior = "RejectMfaByFederatedIdp"
)

func PossibleValuesForInternalDomainFederationFederatedIdpMfaBehavior() []string {
	return []string{
		string(InternalDomainFederationFederatedIdpMfaBehavioracceptIfMfaDoneByFederatedIdp),
		string(InternalDomainFederationFederatedIdpMfaBehaviorenforceMfaByFederatedIdp),
		string(InternalDomainFederationFederatedIdpMfaBehaviorrejectMfaByFederatedIdp),
	}
}

func parseInternalDomainFederationFederatedIdpMfaBehavior(input string) (*InternalDomainFederationFederatedIdpMfaBehavior, error) {
	vals := map[string]InternalDomainFederationFederatedIdpMfaBehavior{
		"acceptifmfadonebyfederatedidp": InternalDomainFederationFederatedIdpMfaBehavioracceptIfMfaDoneByFederatedIdp,
		"enforcemfabyfederatedidp":      InternalDomainFederationFederatedIdpMfaBehaviorenforceMfaByFederatedIdp,
		"rejectmfabyfederatedidp":       InternalDomainFederationFederatedIdpMfaBehaviorrejectMfaByFederatedIdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationFederatedIdpMfaBehavior(input)
	return &out, nil
}
