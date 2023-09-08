package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentDetermination string

const (
	SecurityIncidentDeterminationapt                       SecurityIncidentDetermination = "Apt"
	SecurityIncidentDeterminationcompromisedAccount        SecurityIncidentDetermination = "CompromisedAccount"
	SecurityIncidentDeterminationconfirmedActivity         SecurityIncidentDetermination = "ConfirmedActivity"
	SecurityIncidentDeterminationlineOfBusinessApplication SecurityIncidentDetermination = "LineOfBusinessApplication"
	SecurityIncidentDeterminationmaliciousUserActivity     SecurityIncidentDetermination = "MaliciousUserActivity"
	SecurityIncidentDeterminationmalware                   SecurityIncidentDetermination = "Malware"
	SecurityIncidentDeterminationmultiStagedAttack         SecurityIncidentDetermination = "MultiStagedAttack"
	SecurityIncidentDeterminationnotEnoughDataToValidate   SecurityIncidentDetermination = "NotEnoughDataToValidate"
	SecurityIncidentDeterminationnotMalicious              SecurityIncidentDetermination = "NotMalicious"
	SecurityIncidentDeterminationother                     SecurityIncidentDetermination = "Other"
	SecurityIncidentDeterminationphishing                  SecurityIncidentDetermination = "Phishing"
	SecurityIncidentDeterminationsecurityPersonnel         SecurityIncidentDetermination = "SecurityPersonnel"
	SecurityIncidentDeterminationsecurityTesting           SecurityIncidentDetermination = "SecurityTesting"
	SecurityIncidentDeterminationunknown                   SecurityIncidentDetermination = "Unknown"
	SecurityIncidentDeterminationunwantedSoftware          SecurityIncidentDetermination = "UnwantedSoftware"
)

func PossibleValuesForSecurityIncidentDetermination() []string {
	return []string{
		string(SecurityIncidentDeterminationapt),
		string(SecurityIncidentDeterminationcompromisedAccount),
		string(SecurityIncidentDeterminationconfirmedActivity),
		string(SecurityIncidentDeterminationlineOfBusinessApplication),
		string(SecurityIncidentDeterminationmaliciousUserActivity),
		string(SecurityIncidentDeterminationmalware),
		string(SecurityIncidentDeterminationmultiStagedAttack),
		string(SecurityIncidentDeterminationnotEnoughDataToValidate),
		string(SecurityIncidentDeterminationnotMalicious),
		string(SecurityIncidentDeterminationother),
		string(SecurityIncidentDeterminationphishing),
		string(SecurityIncidentDeterminationsecurityPersonnel),
		string(SecurityIncidentDeterminationsecurityTesting),
		string(SecurityIncidentDeterminationunknown),
		string(SecurityIncidentDeterminationunwantedSoftware),
	}
}

func parseSecurityIncidentDetermination(input string) (*SecurityIncidentDetermination, error) {
	vals := map[string]SecurityIncidentDetermination{
		"apt":                       SecurityIncidentDeterminationapt,
		"compromisedaccount":        SecurityIncidentDeterminationcompromisedAccount,
		"confirmedactivity":         SecurityIncidentDeterminationconfirmedActivity,
		"lineofbusinessapplication": SecurityIncidentDeterminationlineOfBusinessApplication,
		"malicioususeractivity":     SecurityIncidentDeterminationmaliciousUserActivity,
		"malware":                   SecurityIncidentDeterminationmalware,
		"multistagedattack":         SecurityIncidentDeterminationmultiStagedAttack,
		"notenoughdatatovalidate":   SecurityIncidentDeterminationnotEnoughDataToValidate,
		"notmalicious":              SecurityIncidentDeterminationnotMalicious,
		"other":                     SecurityIncidentDeterminationother,
		"phishing":                  SecurityIncidentDeterminationphishing,
		"securitypersonnel":         SecurityIncidentDeterminationsecurityPersonnel,
		"securitytesting":           SecurityIncidentDeterminationsecurityTesting,
		"unknown":                   SecurityIncidentDeterminationunknown,
		"unwantedsoftware":          SecurityIncidentDeterminationunwantedSoftware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentDetermination(input)
	return &out, nil
}
