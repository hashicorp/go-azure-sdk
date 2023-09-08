package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionRiskType string

const (
	RiskDetectionRiskTypeadminConfirmedUserCompromised                RiskDetectionRiskType = "AdminConfirmedUserCompromised"
	RiskDetectionRiskTypeanonymizedIPAddress                          RiskDetectionRiskType = "AnonymizedIPAddress"
	RiskDetectionRiskTypegeneric                                      RiskDetectionRiskType = "Generic"
	RiskDetectionRiskTypeinvestigationsThreatIntelligence             RiskDetectionRiskType = "InvestigationsThreatIntelligence"
	RiskDetectionRiskTypeinvestigationsThreatIntelligenceSigninLinked RiskDetectionRiskType = "InvestigationsThreatIntelligenceSigninLinked"
	RiskDetectionRiskTypeleakedCredentials                            RiskDetectionRiskType = "LeakedCredentials"
	RiskDetectionRiskTypemaliciousIPAddress                           RiskDetectionRiskType = "MaliciousIPAddress"
	RiskDetectionRiskTypemaliciousIPAddressValidCredentialsBlockedIP  RiskDetectionRiskType = "MaliciousIPAddressValidCredentialsBlockedIP"
	RiskDetectionRiskTypemalwareInfectedIPAddress                     RiskDetectionRiskType = "MalwareInfectedIPAddress"
	RiskDetectionRiskTypemcasImpossibleTravel                         RiskDetectionRiskType = "McasImpossibleTravel"
	RiskDetectionRiskTypemcasSuspiciousInboxManipulationRules         RiskDetectionRiskType = "McasSuspiciousInboxManipulationRules"
	RiskDetectionRiskTypesuspiciousIPAddress                          RiskDetectionRiskType = "SuspiciousIPAddress"
	RiskDetectionRiskTypeunfamiliarFeatures                           RiskDetectionRiskType = "UnfamiliarFeatures"
	RiskDetectionRiskTypeunlikelyTravel                               RiskDetectionRiskType = "UnlikelyTravel"
)

func PossibleValuesForRiskDetectionRiskType() []string {
	return []string{
		string(RiskDetectionRiskTypeadminConfirmedUserCompromised),
		string(RiskDetectionRiskTypeanonymizedIPAddress),
		string(RiskDetectionRiskTypegeneric),
		string(RiskDetectionRiskTypeinvestigationsThreatIntelligence),
		string(RiskDetectionRiskTypeinvestigationsThreatIntelligenceSigninLinked),
		string(RiskDetectionRiskTypeleakedCredentials),
		string(RiskDetectionRiskTypemaliciousIPAddress),
		string(RiskDetectionRiskTypemaliciousIPAddressValidCredentialsBlockedIP),
		string(RiskDetectionRiskTypemalwareInfectedIPAddress),
		string(RiskDetectionRiskTypemcasImpossibleTravel),
		string(RiskDetectionRiskTypemcasSuspiciousInboxManipulationRules),
		string(RiskDetectionRiskTypesuspiciousIPAddress),
		string(RiskDetectionRiskTypeunfamiliarFeatures),
		string(RiskDetectionRiskTypeunlikelyTravel),
	}
}

func parseRiskDetectionRiskType(input string) (*RiskDetectionRiskType, error) {
	vals := map[string]RiskDetectionRiskType{
		"adminconfirmedusercompromised":                RiskDetectionRiskTypeadminConfirmedUserCompromised,
		"anonymizedipaddress":                          RiskDetectionRiskTypeanonymizedIPAddress,
		"generic":                                      RiskDetectionRiskTypegeneric,
		"investigationsthreatintelligence":             RiskDetectionRiskTypeinvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": RiskDetectionRiskTypeinvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            RiskDetectionRiskTypeleakedCredentials,
		"maliciousipaddress":                           RiskDetectionRiskTypemaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  RiskDetectionRiskTypemaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     RiskDetectionRiskTypemalwareInfectedIPAddress,
		"mcasimpossibletravel":                         RiskDetectionRiskTypemcasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         RiskDetectionRiskTypemcasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          RiskDetectionRiskTypesuspiciousIPAddress,
		"unfamiliarfeatures":                           RiskDetectionRiskTypeunfamiliarFeatures,
		"unlikelytravel":                               RiskDetectionRiskTypeunlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionRiskType(input)
	return &out, nil
}
