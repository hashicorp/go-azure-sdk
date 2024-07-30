package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivityEventTypes string

const (
	RiskUserActivityEventTypes_AdminConfirmedUserCompromised                RiskUserActivityEventTypes = "adminConfirmedUserCompromised"
	RiskUserActivityEventTypes_AnonymizedIPAddress                          RiskUserActivityEventTypes = "anonymizedIPAddress"
	RiskUserActivityEventTypes_Generic                                      RiskUserActivityEventTypes = "generic"
	RiskUserActivityEventTypes_InvestigationsThreatIntelligence             RiskUserActivityEventTypes = "investigationsThreatIntelligence"
	RiskUserActivityEventTypes_InvestigationsThreatIntelligenceSigninLinked RiskUserActivityEventTypes = "investigationsThreatIntelligenceSigninLinked"
	RiskUserActivityEventTypes_LeakedCredentials                            RiskUserActivityEventTypes = "leakedCredentials"
	RiskUserActivityEventTypes_MaliciousIPAddress                           RiskUserActivityEventTypes = "maliciousIPAddress"
	RiskUserActivityEventTypes_MaliciousIPAddressValidCredentialsBlockedIP  RiskUserActivityEventTypes = "maliciousIPAddressValidCredentialsBlockedIP"
	RiskUserActivityEventTypes_MalwareInfectedIPAddress                     RiskUserActivityEventTypes = "malwareInfectedIPAddress"
	RiskUserActivityEventTypes_McasImpossibleTravel                         RiskUserActivityEventTypes = "mcasImpossibleTravel"
	RiskUserActivityEventTypes_McasSuspiciousInboxManipulationRules         RiskUserActivityEventTypes = "mcasSuspiciousInboxManipulationRules"
	RiskUserActivityEventTypes_SuspiciousIPAddress                          RiskUserActivityEventTypes = "suspiciousIPAddress"
	RiskUserActivityEventTypes_UnfamiliarFeatures                           RiskUserActivityEventTypes = "unfamiliarFeatures"
	RiskUserActivityEventTypes_UnlikelyTravel                               RiskUserActivityEventTypes = "unlikelyTravel"
)

func PossibleValuesForRiskUserActivityEventTypes() []string {
	return []string{
		string(RiskUserActivityEventTypes_AdminConfirmedUserCompromised),
		string(RiskUserActivityEventTypes_AnonymizedIPAddress),
		string(RiskUserActivityEventTypes_Generic),
		string(RiskUserActivityEventTypes_InvestigationsThreatIntelligence),
		string(RiskUserActivityEventTypes_InvestigationsThreatIntelligenceSigninLinked),
		string(RiskUserActivityEventTypes_LeakedCredentials),
		string(RiskUserActivityEventTypes_MaliciousIPAddress),
		string(RiskUserActivityEventTypes_MaliciousIPAddressValidCredentialsBlockedIP),
		string(RiskUserActivityEventTypes_MalwareInfectedIPAddress),
		string(RiskUserActivityEventTypes_McasImpossibleTravel),
		string(RiskUserActivityEventTypes_McasSuspiciousInboxManipulationRules),
		string(RiskUserActivityEventTypes_SuspiciousIPAddress),
		string(RiskUserActivityEventTypes_UnfamiliarFeatures),
		string(RiskUserActivityEventTypes_UnlikelyTravel),
	}
}

func (s *RiskUserActivityEventTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskUserActivityEventTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskUserActivityEventTypes(input string) (*RiskUserActivityEventTypes, error) {
	vals := map[string]RiskUserActivityEventTypes{
		"adminconfirmedusercompromised":                RiskUserActivityEventTypes_AdminConfirmedUserCompromised,
		"anonymizedipaddress":                          RiskUserActivityEventTypes_AnonymizedIPAddress,
		"generic":                                      RiskUserActivityEventTypes_Generic,
		"investigationsthreatintelligence":             RiskUserActivityEventTypes_InvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": RiskUserActivityEventTypes_InvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            RiskUserActivityEventTypes_LeakedCredentials,
		"maliciousipaddress":                           RiskUserActivityEventTypes_MaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  RiskUserActivityEventTypes_MaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     RiskUserActivityEventTypes_MalwareInfectedIPAddress,
		"mcasimpossibletravel":                         RiskUserActivityEventTypes_McasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         RiskUserActivityEventTypes_McasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          RiskUserActivityEventTypes_SuspiciousIPAddress,
		"unfamiliarfeatures":                           RiskUserActivityEventTypes_UnfamiliarFeatures,
		"unlikelytravel":                               RiskUserActivityEventTypes_UnlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskUserActivityEventTypes(input)
	return &out, nil
}
