package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityToolAwsUserAdministratorFindingSecurityTools string

const (
	SecurityToolAwsUserAdministratorFindingSecurityTools_CloudTrail  SecurityToolAwsUserAdministratorFindingSecurityTools = "cloudTrail"
	SecurityToolAwsUserAdministratorFindingSecurityTools_Detective   SecurityToolAwsUserAdministratorFindingSecurityTools = "detective"
	SecurityToolAwsUserAdministratorFindingSecurityTools_GuardDuty   SecurityToolAwsUserAdministratorFindingSecurityTools = "guardDuty"
	SecurityToolAwsUserAdministratorFindingSecurityTools_Inspector   SecurityToolAwsUserAdministratorFindingSecurityTools = "inspector"
	SecurityToolAwsUserAdministratorFindingSecurityTools_Macie       SecurityToolAwsUserAdministratorFindingSecurityTools = "macie"
	SecurityToolAwsUserAdministratorFindingSecurityTools_SecurityHub SecurityToolAwsUserAdministratorFindingSecurityTools = "securityHub"
	SecurityToolAwsUserAdministratorFindingSecurityTools_WafShield   SecurityToolAwsUserAdministratorFindingSecurityTools = "wafShield"
)

func PossibleValuesForSecurityToolAwsUserAdministratorFindingSecurityTools() []string {
	return []string{
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_CloudTrail),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_Detective),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_GuardDuty),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_Inspector),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_Macie),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_SecurityHub),
		string(SecurityToolAwsUserAdministratorFindingSecurityTools_WafShield),
	}
}

func (s *SecurityToolAwsUserAdministratorFindingSecurityTools) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityToolAwsUserAdministratorFindingSecurityTools(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityToolAwsUserAdministratorFindingSecurityTools(input string) (*SecurityToolAwsUserAdministratorFindingSecurityTools, error) {
	vals := map[string]SecurityToolAwsUserAdministratorFindingSecurityTools{
		"cloudtrail":  SecurityToolAwsUserAdministratorFindingSecurityTools_CloudTrail,
		"detective":   SecurityToolAwsUserAdministratorFindingSecurityTools_Detective,
		"guardduty":   SecurityToolAwsUserAdministratorFindingSecurityTools_GuardDuty,
		"inspector":   SecurityToolAwsUserAdministratorFindingSecurityTools_Inspector,
		"macie":       SecurityToolAwsUserAdministratorFindingSecurityTools_Macie,
		"securityhub": SecurityToolAwsUserAdministratorFindingSecurityTools_SecurityHub,
		"wafshield":   SecurityToolAwsUserAdministratorFindingSecurityTools_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityToolAwsUserAdministratorFindingSecurityTools(input)
	return &out, nil
}
