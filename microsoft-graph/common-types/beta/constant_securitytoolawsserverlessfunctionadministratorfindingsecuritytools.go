package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools string

const (
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_CloudTrail  SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "cloudTrail"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Detective   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "detective"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_GuardDuty   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "guardDuty"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Inspector   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "inspector"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Macie       SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "macie"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_SecurityHub SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "securityHub"
	SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_WafShield   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools = "wafShield"
)

func PossibleValuesForSecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools() []string {
	return []string{
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_CloudTrail),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Detective),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_GuardDuty),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Inspector),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Macie),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_SecurityHub),
		string(SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_WafShield),
	}
}

func (s *SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools(input string) (*SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools, error) {
	vals := map[string]SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools{
		"cloudtrail":  SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_CloudTrail,
		"detective":   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Detective,
		"guardduty":   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_GuardDuty,
		"inspector":   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Inspector,
		"macie":       SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_Macie,
		"securityhub": SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_SecurityHub,
		"wafshield":   SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityToolAwsServerlessFunctionAdministratorFindingSecurityTools(input)
	return &out, nil
}
