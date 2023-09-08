package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertDetermination string

const (
	SecurityAlertDeterminationapt                       SecurityAlertDetermination = "Apt"
	SecurityAlertDeterminationcompromisedAccount        SecurityAlertDetermination = "CompromisedAccount"
	SecurityAlertDeterminationconfirmedActivity         SecurityAlertDetermination = "ConfirmedActivity"
	SecurityAlertDeterminationlineOfBusinessApplication SecurityAlertDetermination = "LineOfBusinessApplication"
	SecurityAlertDeterminationmaliciousUserActivity     SecurityAlertDetermination = "MaliciousUserActivity"
	SecurityAlertDeterminationmalware                   SecurityAlertDetermination = "Malware"
	SecurityAlertDeterminationmultiStagedAttack         SecurityAlertDetermination = "MultiStagedAttack"
	SecurityAlertDeterminationnotEnoughDataToValidate   SecurityAlertDetermination = "NotEnoughDataToValidate"
	SecurityAlertDeterminationnotMalicious              SecurityAlertDetermination = "NotMalicious"
	SecurityAlertDeterminationother                     SecurityAlertDetermination = "Other"
	SecurityAlertDeterminationphishing                  SecurityAlertDetermination = "Phishing"
	SecurityAlertDeterminationsecurityPersonnel         SecurityAlertDetermination = "SecurityPersonnel"
	SecurityAlertDeterminationsecurityTesting           SecurityAlertDetermination = "SecurityTesting"
	SecurityAlertDeterminationunknown                   SecurityAlertDetermination = "Unknown"
	SecurityAlertDeterminationunwantedSoftware          SecurityAlertDetermination = "UnwantedSoftware"
)

func PossibleValuesForSecurityAlertDetermination() []string {
	return []string{
		string(SecurityAlertDeterminationapt),
		string(SecurityAlertDeterminationcompromisedAccount),
		string(SecurityAlertDeterminationconfirmedActivity),
		string(SecurityAlertDeterminationlineOfBusinessApplication),
		string(SecurityAlertDeterminationmaliciousUserActivity),
		string(SecurityAlertDeterminationmalware),
		string(SecurityAlertDeterminationmultiStagedAttack),
		string(SecurityAlertDeterminationnotEnoughDataToValidate),
		string(SecurityAlertDeterminationnotMalicious),
		string(SecurityAlertDeterminationother),
		string(SecurityAlertDeterminationphishing),
		string(SecurityAlertDeterminationsecurityPersonnel),
		string(SecurityAlertDeterminationsecurityTesting),
		string(SecurityAlertDeterminationunknown),
		string(SecurityAlertDeterminationunwantedSoftware),
	}
}

func parseSecurityAlertDetermination(input string) (*SecurityAlertDetermination, error) {
	vals := map[string]SecurityAlertDetermination{
		"apt":                       SecurityAlertDeterminationapt,
		"compromisedaccount":        SecurityAlertDeterminationcompromisedAccount,
		"confirmedactivity":         SecurityAlertDeterminationconfirmedActivity,
		"lineofbusinessapplication": SecurityAlertDeterminationlineOfBusinessApplication,
		"malicioususeractivity":     SecurityAlertDeterminationmaliciousUserActivity,
		"malware":                   SecurityAlertDeterminationmalware,
		"multistagedattack":         SecurityAlertDeterminationmultiStagedAttack,
		"notenoughdatatovalidate":   SecurityAlertDeterminationnotEnoughDataToValidate,
		"notmalicious":              SecurityAlertDeterminationnotMalicious,
		"other":                     SecurityAlertDeterminationother,
		"phishing":                  SecurityAlertDeterminationphishing,
		"securitypersonnel":         SecurityAlertDeterminationsecurityPersonnel,
		"securitytesting":           SecurityAlertDeterminationsecurityTesting,
		"unknown":                   SecurityAlertDeterminationunknown,
		"unwantedsoftware":          SecurityAlertDeterminationunwantedSoftware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertDetermination(input)
	return &out, nil
}
