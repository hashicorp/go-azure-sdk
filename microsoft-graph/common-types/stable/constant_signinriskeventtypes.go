package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskEventTypes string

const (
	SignInRiskEventTypes_AdminConfirmedUserCompromised                SignInRiskEventTypes = "adminConfirmedUserCompromised"
	SignInRiskEventTypes_AnonymizedIPAddress                          SignInRiskEventTypes = "anonymizedIPAddress"
	SignInRiskEventTypes_Generic                                      SignInRiskEventTypes = "generic"
	SignInRiskEventTypes_InvestigationsThreatIntelligence             SignInRiskEventTypes = "investigationsThreatIntelligence"
	SignInRiskEventTypes_InvestigationsThreatIntelligenceSigninLinked SignInRiskEventTypes = "investigationsThreatIntelligenceSigninLinked"
	SignInRiskEventTypes_LeakedCredentials                            SignInRiskEventTypes = "leakedCredentials"
	SignInRiskEventTypes_MaliciousIPAddress                           SignInRiskEventTypes = "maliciousIPAddress"
	SignInRiskEventTypes_MaliciousIPAddressValidCredentialsBlockedIP  SignInRiskEventTypes = "maliciousIPAddressValidCredentialsBlockedIP"
	SignInRiskEventTypes_MalwareInfectedIPAddress                     SignInRiskEventTypes = "malwareInfectedIPAddress"
	SignInRiskEventTypes_McasImpossibleTravel                         SignInRiskEventTypes = "mcasImpossibleTravel"
	SignInRiskEventTypes_McasSuspiciousInboxManipulationRules         SignInRiskEventTypes = "mcasSuspiciousInboxManipulationRules"
	SignInRiskEventTypes_SuspiciousIPAddress                          SignInRiskEventTypes = "suspiciousIPAddress"
	SignInRiskEventTypes_UnfamiliarFeatures                           SignInRiskEventTypes = "unfamiliarFeatures"
	SignInRiskEventTypes_UnlikelyTravel                               SignInRiskEventTypes = "unlikelyTravel"
)

func PossibleValuesForSignInRiskEventTypes() []string {
	return []string{
		string(SignInRiskEventTypes_AdminConfirmedUserCompromised),
		string(SignInRiskEventTypes_AnonymizedIPAddress),
		string(SignInRiskEventTypes_Generic),
		string(SignInRiskEventTypes_InvestigationsThreatIntelligence),
		string(SignInRiskEventTypes_InvestigationsThreatIntelligenceSigninLinked),
		string(SignInRiskEventTypes_LeakedCredentials),
		string(SignInRiskEventTypes_MaliciousIPAddress),
		string(SignInRiskEventTypes_MaliciousIPAddressValidCredentialsBlockedIP),
		string(SignInRiskEventTypes_MalwareInfectedIPAddress),
		string(SignInRiskEventTypes_McasImpossibleTravel),
		string(SignInRiskEventTypes_McasSuspiciousInboxManipulationRules),
		string(SignInRiskEventTypes_SuspiciousIPAddress),
		string(SignInRiskEventTypes_UnfamiliarFeatures),
		string(SignInRiskEventTypes_UnlikelyTravel),
	}
}

func (s *SignInRiskEventTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInRiskEventTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInRiskEventTypes(input string) (*SignInRiskEventTypes, error) {
	vals := map[string]SignInRiskEventTypes{
		"adminconfirmedusercompromised":                SignInRiskEventTypes_AdminConfirmedUserCompromised,
		"anonymizedipaddress":                          SignInRiskEventTypes_AnonymizedIPAddress,
		"generic":                                      SignInRiskEventTypes_Generic,
		"investigationsthreatintelligence":             SignInRiskEventTypes_InvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": SignInRiskEventTypes_InvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            SignInRiskEventTypes_LeakedCredentials,
		"maliciousipaddress":                           SignInRiskEventTypes_MaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  SignInRiskEventTypes_MaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     SignInRiskEventTypes_MalwareInfectedIPAddress,
		"mcasimpossibletravel":                         SignInRiskEventTypes_McasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         SignInRiskEventTypes_McasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          SignInRiskEventTypes_SuspiciousIPAddress,
		"unfamiliarfeatures":                           SignInRiskEventTypes_UnfamiliarFeatures,
		"unlikelytravel":                               SignInRiskEventTypes_UnlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskEventTypes(input)
	return &out, nil
}
