package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivityEventTypes string

const (
	RiskUserActivityEventTypesadminConfirmedUserCompromised                RiskUserActivityEventTypes = "AdminConfirmedUserCompromised"
	RiskUserActivityEventTypesanonymizedIPAddress                          RiskUserActivityEventTypes = "AnonymizedIPAddress"
	RiskUserActivityEventTypesgeneric                                      RiskUserActivityEventTypes = "Generic"
	RiskUserActivityEventTypesinvestigationsThreatIntelligence             RiskUserActivityEventTypes = "InvestigationsThreatIntelligence"
	RiskUserActivityEventTypesinvestigationsThreatIntelligenceSigninLinked RiskUserActivityEventTypes = "InvestigationsThreatIntelligenceSigninLinked"
	RiskUserActivityEventTypesleakedCredentials                            RiskUserActivityEventTypes = "LeakedCredentials"
	RiskUserActivityEventTypesmaliciousIPAddress                           RiskUserActivityEventTypes = "MaliciousIPAddress"
	RiskUserActivityEventTypesmaliciousIPAddressValidCredentialsBlockedIP  RiskUserActivityEventTypes = "MaliciousIPAddressValidCredentialsBlockedIP"
	RiskUserActivityEventTypesmalwareInfectedIPAddress                     RiskUserActivityEventTypes = "MalwareInfectedIPAddress"
	RiskUserActivityEventTypesmcasImpossibleTravel                         RiskUserActivityEventTypes = "McasImpossibleTravel"
	RiskUserActivityEventTypesmcasSuspiciousInboxManipulationRules         RiskUserActivityEventTypes = "McasSuspiciousInboxManipulationRules"
	RiskUserActivityEventTypessuspiciousIPAddress                          RiskUserActivityEventTypes = "SuspiciousIPAddress"
	RiskUserActivityEventTypesunfamiliarFeatures                           RiskUserActivityEventTypes = "UnfamiliarFeatures"
	RiskUserActivityEventTypesunlikelyTravel                               RiskUserActivityEventTypes = "UnlikelyTravel"
)

func PossibleValuesForRiskUserActivityEventTypes() []string {
	return []string{
		string(RiskUserActivityEventTypesadminConfirmedUserCompromised),
		string(RiskUserActivityEventTypesanonymizedIPAddress),
		string(RiskUserActivityEventTypesgeneric),
		string(RiskUserActivityEventTypesinvestigationsThreatIntelligence),
		string(RiskUserActivityEventTypesinvestigationsThreatIntelligenceSigninLinked),
		string(RiskUserActivityEventTypesleakedCredentials),
		string(RiskUserActivityEventTypesmaliciousIPAddress),
		string(RiskUserActivityEventTypesmaliciousIPAddressValidCredentialsBlockedIP),
		string(RiskUserActivityEventTypesmalwareInfectedIPAddress),
		string(RiskUserActivityEventTypesmcasImpossibleTravel),
		string(RiskUserActivityEventTypesmcasSuspiciousInboxManipulationRules),
		string(RiskUserActivityEventTypessuspiciousIPAddress),
		string(RiskUserActivityEventTypesunfamiliarFeatures),
		string(RiskUserActivityEventTypesunlikelyTravel),
	}
}

func parseRiskUserActivityEventTypes(input string) (*RiskUserActivityEventTypes, error) {
	vals := map[string]RiskUserActivityEventTypes{
		"adminconfirmedusercompromised":                RiskUserActivityEventTypesadminConfirmedUserCompromised,
		"anonymizedipaddress":                          RiskUserActivityEventTypesanonymizedIPAddress,
		"generic":                                      RiskUserActivityEventTypesgeneric,
		"investigationsthreatintelligence":             RiskUserActivityEventTypesinvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": RiskUserActivityEventTypesinvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            RiskUserActivityEventTypesleakedCredentials,
		"maliciousipaddress":                           RiskUserActivityEventTypesmaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  RiskUserActivityEventTypesmaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     RiskUserActivityEventTypesmalwareInfectedIPAddress,
		"mcasimpossibletravel":                         RiskUserActivityEventTypesmcasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         RiskUserActivityEventTypesmcasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          RiskUserActivityEventTypessuspiciousIPAddress,
		"unfamiliarfeatures":                           RiskUserActivityEventTypesunfamiliarFeatures,
		"unlikelytravel":                               RiskUserActivityEventTypesunlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskUserActivityEventTypes(input)
	return &out, nil
}
