package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentDetermination string

const (
	SecurityIncidentDetermination_Apt                       SecurityIncidentDetermination = "apt"
	SecurityIncidentDetermination_CompromisedAccount        SecurityIncidentDetermination = "compromisedAccount"
	SecurityIncidentDetermination_ConfirmedActivity         SecurityIncidentDetermination = "confirmedActivity"
	SecurityIncidentDetermination_LineOfBusinessApplication SecurityIncidentDetermination = "lineOfBusinessApplication"
	SecurityIncidentDetermination_MaliciousUserActivity     SecurityIncidentDetermination = "maliciousUserActivity"
	SecurityIncidentDetermination_Malware                   SecurityIncidentDetermination = "malware"
	SecurityIncidentDetermination_MultiStagedAttack         SecurityIncidentDetermination = "multiStagedAttack"
	SecurityIncidentDetermination_NotEnoughDataToValidate   SecurityIncidentDetermination = "notEnoughDataToValidate"
	SecurityIncidentDetermination_NotMalicious              SecurityIncidentDetermination = "notMalicious"
	SecurityIncidentDetermination_Other                     SecurityIncidentDetermination = "other"
	SecurityIncidentDetermination_Phishing                  SecurityIncidentDetermination = "phishing"
	SecurityIncidentDetermination_SecurityPersonnel         SecurityIncidentDetermination = "securityPersonnel"
	SecurityIncidentDetermination_SecurityTesting           SecurityIncidentDetermination = "securityTesting"
	SecurityIncidentDetermination_Unknown                   SecurityIncidentDetermination = "unknown"
	SecurityIncidentDetermination_UnwantedSoftware          SecurityIncidentDetermination = "unwantedSoftware"
)

func PossibleValuesForSecurityIncidentDetermination() []string {
	return []string{
		string(SecurityIncidentDetermination_Apt),
		string(SecurityIncidentDetermination_CompromisedAccount),
		string(SecurityIncidentDetermination_ConfirmedActivity),
		string(SecurityIncidentDetermination_LineOfBusinessApplication),
		string(SecurityIncidentDetermination_MaliciousUserActivity),
		string(SecurityIncidentDetermination_Malware),
		string(SecurityIncidentDetermination_MultiStagedAttack),
		string(SecurityIncidentDetermination_NotEnoughDataToValidate),
		string(SecurityIncidentDetermination_NotMalicious),
		string(SecurityIncidentDetermination_Other),
		string(SecurityIncidentDetermination_Phishing),
		string(SecurityIncidentDetermination_SecurityPersonnel),
		string(SecurityIncidentDetermination_SecurityTesting),
		string(SecurityIncidentDetermination_Unknown),
		string(SecurityIncidentDetermination_UnwantedSoftware),
	}
}

func (s *SecurityIncidentDetermination) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIncidentDetermination(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIncidentDetermination(input string) (*SecurityIncidentDetermination, error) {
	vals := map[string]SecurityIncidentDetermination{
		"apt":                       SecurityIncidentDetermination_Apt,
		"compromisedaccount":        SecurityIncidentDetermination_CompromisedAccount,
		"confirmedactivity":         SecurityIncidentDetermination_ConfirmedActivity,
		"lineofbusinessapplication": SecurityIncidentDetermination_LineOfBusinessApplication,
		"malicioususeractivity":     SecurityIncidentDetermination_MaliciousUserActivity,
		"malware":                   SecurityIncidentDetermination_Malware,
		"multistagedattack":         SecurityIncidentDetermination_MultiStagedAttack,
		"notenoughdatatovalidate":   SecurityIncidentDetermination_NotEnoughDataToValidate,
		"notmalicious":              SecurityIncidentDetermination_NotMalicious,
		"other":                     SecurityIncidentDetermination_Other,
		"phishing":                  SecurityIncidentDetermination_Phishing,
		"securitypersonnel":         SecurityIncidentDetermination_SecurityPersonnel,
		"securitytesting":           SecurityIncidentDetermination_SecurityTesting,
		"unknown":                   SecurityIncidentDetermination_Unknown,
		"unwantedsoftware":          SecurityIncidentDetermination_UnwantedSoftware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentDetermination(input)
	return &out, nil
}
