package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskType string

const (
	RiskDetectionRiskType_AdminConfirmedUserCompromised                RiskDetectionRiskType = "adminConfirmedUserCompromised"
	RiskDetectionRiskType_AnonymizedIPAddress                          RiskDetectionRiskType = "anonymizedIPAddress"
	RiskDetectionRiskType_Generic                                      RiskDetectionRiskType = "generic"
	RiskDetectionRiskType_InvestigationsThreatIntelligence             RiskDetectionRiskType = "investigationsThreatIntelligence"
	RiskDetectionRiskType_InvestigationsThreatIntelligenceSigninLinked RiskDetectionRiskType = "investigationsThreatIntelligenceSigninLinked"
	RiskDetectionRiskType_LeakedCredentials                            RiskDetectionRiskType = "leakedCredentials"
	RiskDetectionRiskType_MaliciousIPAddress                           RiskDetectionRiskType = "maliciousIPAddress"
	RiskDetectionRiskType_MaliciousIPAddressValidCredentialsBlockedIP  RiskDetectionRiskType = "maliciousIPAddressValidCredentialsBlockedIP"
	RiskDetectionRiskType_MalwareInfectedIPAddress                     RiskDetectionRiskType = "malwareInfectedIPAddress"
	RiskDetectionRiskType_McasImpossibleTravel                         RiskDetectionRiskType = "mcasImpossibleTravel"
	RiskDetectionRiskType_McasSuspiciousInboxManipulationRules         RiskDetectionRiskType = "mcasSuspiciousInboxManipulationRules"
	RiskDetectionRiskType_SuspiciousIPAddress                          RiskDetectionRiskType = "suspiciousIPAddress"
	RiskDetectionRiskType_UnfamiliarFeatures                           RiskDetectionRiskType = "unfamiliarFeatures"
	RiskDetectionRiskType_UnlikelyTravel                               RiskDetectionRiskType = "unlikelyTravel"
)

func PossibleValuesForRiskDetectionRiskType() []string {
	return []string{
		string(RiskDetectionRiskType_AdminConfirmedUserCompromised),
		string(RiskDetectionRiskType_AnonymizedIPAddress),
		string(RiskDetectionRiskType_Generic),
		string(RiskDetectionRiskType_InvestigationsThreatIntelligence),
		string(RiskDetectionRiskType_InvestigationsThreatIntelligenceSigninLinked),
		string(RiskDetectionRiskType_LeakedCredentials),
		string(RiskDetectionRiskType_MaliciousIPAddress),
		string(RiskDetectionRiskType_MaliciousIPAddressValidCredentialsBlockedIP),
		string(RiskDetectionRiskType_MalwareInfectedIPAddress),
		string(RiskDetectionRiskType_McasImpossibleTravel),
		string(RiskDetectionRiskType_McasSuspiciousInboxManipulationRules),
		string(RiskDetectionRiskType_SuspiciousIPAddress),
		string(RiskDetectionRiskType_UnfamiliarFeatures),
		string(RiskDetectionRiskType_UnlikelyTravel),
	}
}

func (s *RiskDetectionRiskType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionRiskType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionRiskType(input string) (*RiskDetectionRiskType, error) {
	vals := map[string]RiskDetectionRiskType{
		"adminconfirmedusercompromised":                RiskDetectionRiskType_AdminConfirmedUserCompromised,
		"anonymizedipaddress":                          RiskDetectionRiskType_AnonymizedIPAddress,
		"generic":                                      RiskDetectionRiskType_Generic,
		"investigationsthreatintelligence":             RiskDetectionRiskType_InvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": RiskDetectionRiskType_InvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            RiskDetectionRiskType_LeakedCredentials,
		"maliciousipaddress":                           RiskDetectionRiskType_MaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  RiskDetectionRiskType_MaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     RiskDetectionRiskType_MalwareInfectedIPAddress,
		"mcasimpossibletravel":                         RiskDetectionRiskType_McasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         RiskDetectionRiskType_McasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          RiskDetectionRiskType_SuspiciousIPAddress,
		"unfamiliarfeatures":                           RiskDetectionRiskType_UnfamiliarFeatures,
		"unlikelytravel":                               RiskDetectionRiskType_UnlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskType(input)
	return &out, nil
}
