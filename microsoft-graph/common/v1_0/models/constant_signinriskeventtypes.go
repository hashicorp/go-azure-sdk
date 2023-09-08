package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskEventTypes string

const (
	SignInRiskEventTypesadminConfirmedUserCompromised                SignInRiskEventTypes = "AdminConfirmedUserCompromised"
	SignInRiskEventTypesanonymizedIPAddress                          SignInRiskEventTypes = "AnonymizedIPAddress"
	SignInRiskEventTypesgeneric                                      SignInRiskEventTypes = "Generic"
	SignInRiskEventTypesinvestigationsThreatIntelligence             SignInRiskEventTypes = "InvestigationsThreatIntelligence"
	SignInRiskEventTypesinvestigationsThreatIntelligenceSigninLinked SignInRiskEventTypes = "InvestigationsThreatIntelligenceSigninLinked"
	SignInRiskEventTypesleakedCredentials                            SignInRiskEventTypes = "LeakedCredentials"
	SignInRiskEventTypesmaliciousIPAddress                           SignInRiskEventTypes = "MaliciousIPAddress"
	SignInRiskEventTypesmaliciousIPAddressValidCredentialsBlockedIP  SignInRiskEventTypes = "MaliciousIPAddressValidCredentialsBlockedIP"
	SignInRiskEventTypesmalwareInfectedIPAddress                     SignInRiskEventTypes = "MalwareInfectedIPAddress"
	SignInRiskEventTypesmcasImpossibleTravel                         SignInRiskEventTypes = "McasImpossibleTravel"
	SignInRiskEventTypesmcasSuspiciousInboxManipulationRules         SignInRiskEventTypes = "McasSuspiciousInboxManipulationRules"
	SignInRiskEventTypessuspiciousIPAddress                          SignInRiskEventTypes = "SuspiciousIPAddress"
	SignInRiskEventTypesunfamiliarFeatures                           SignInRiskEventTypes = "UnfamiliarFeatures"
	SignInRiskEventTypesunlikelyTravel                               SignInRiskEventTypes = "UnlikelyTravel"
)

func PossibleValuesForSignInRiskEventTypes() []string {
	return []string{
		string(SignInRiskEventTypesadminConfirmedUserCompromised),
		string(SignInRiskEventTypesanonymizedIPAddress),
		string(SignInRiskEventTypesgeneric),
		string(SignInRiskEventTypesinvestigationsThreatIntelligence),
		string(SignInRiskEventTypesinvestigationsThreatIntelligenceSigninLinked),
		string(SignInRiskEventTypesleakedCredentials),
		string(SignInRiskEventTypesmaliciousIPAddress),
		string(SignInRiskEventTypesmaliciousIPAddressValidCredentialsBlockedIP),
		string(SignInRiskEventTypesmalwareInfectedIPAddress),
		string(SignInRiskEventTypesmcasImpossibleTravel),
		string(SignInRiskEventTypesmcasSuspiciousInboxManipulationRules),
		string(SignInRiskEventTypessuspiciousIPAddress),
		string(SignInRiskEventTypesunfamiliarFeatures),
		string(SignInRiskEventTypesunlikelyTravel),
	}
}

func parseSignInRiskEventTypes(input string) (*SignInRiskEventTypes, error) {
	vals := map[string]SignInRiskEventTypes{
		"adminconfirmedusercompromised":                SignInRiskEventTypesadminConfirmedUserCompromised,
		"anonymizedipaddress":                          SignInRiskEventTypesanonymizedIPAddress,
		"generic":                                      SignInRiskEventTypesgeneric,
		"investigationsthreatintelligence":             SignInRiskEventTypesinvestigationsThreatIntelligence,
		"investigationsthreatintelligencesigninlinked": SignInRiskEventTypesinvestigationsThreatIntelligenceSigninLinked,
		"leakedcredentials":                            SignInRiskEventTypesleakedCredentials,
		"maliciousipaddress":                           SignInRiskEventTypesmaliciousIPAddress,
		"maliciousipaddressvalidcredentialsblockedip":  SignInRiskEventTypesmaliciousIPAddressValidCredentialsBlockedIP,
		"malwareinfectedipaddress":                     SignInRiskEventTypesmalwareInfectedIPAddress,
		"mcasimpossibletravel":                         SignInRiskEventTypesmcasImpossibleTravel,
		"mcassuspiciousinboxmanipulationrules":         SignInRiskEventTypesmcasSuspiciousInboxManipulationRules,
		"suspiciousipaddress":                          SignInRiskEventTypessuspiciousIPAddress,
		"unfamiliarfeatures":                           SignInRiskEventTypesunfamiliarFeatures,
		"unlikelytravel":                               SignInRiskEventTypesunlikelyTravel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskEventTypes(input)
	return &out, nil
}
