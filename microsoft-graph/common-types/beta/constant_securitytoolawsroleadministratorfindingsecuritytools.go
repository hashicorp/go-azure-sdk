package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityToolAwsRoleAdministratorFindingSecurityTools string

const (
	SecurityToolAwsRoleAdministratorFindingSecurityTools_CloudTrail  SecurityToolAwsRoleAdministratorFindingSecurityTools = "cloudTrail"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_Detective   SecurityToolAwsRoleAdministratorFindingSecurityTools = "detective"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_GuardDuty   SecurityToolAwsRoleAdministratorFindingSecurityTools = "guardDuty"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_Inspector   SecurityToolAwsRoleAdministratorFindingSecurityTools = "inspector"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_Macie       SecurityToolAwsRoleAdministratorFindingSecurityTools = "macie"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_SecurityHub SecurityToolAwsRoleAdministratorFindingSecurityTools = "securityHub"
	SecurityToolAwsRoleAdministratorFindingSecurityTools_WafShield   SecurityToolAwsRoleAdministratorFindingSecurityTools = "wafShield"
)

func PossibleValuesForSecurityToolAwsRoleAdministratorFindingSecurityTools() []string {
	return []string{
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_CloudTrail),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_Detective),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_GuardDuty),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_Inspector),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_Macie),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_SecurityHub),
		string(SecurityToolAwsRoleAdministratorFindingSecurityTools_WafShield),
	}
}

func (s *SecurityToolAwsRoleAdministratorFindingSecurityTools) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityToolAwsRoleAdministratorFindingSecurityTools(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityToolAwsRoleAdministratorFindingSecurityTools(input string) (*SecurityToolAwsRoleAdministratorFindingSecurityTools, error) {
	vals := map[string]SecurityToolAwsRoleAdministratorFindingSecurityTools{
		"cloudtrail":  SecurityToolAwsRoleAdministratorFindingSecurityTools_CloudTrail,
		"detective":   SecurityToolAwsRoleAdministratorFindingSecurityTools_Detective,
		"guardduty":   SecurityToolAwsRoleAdministratorFindingSecurityTools_GuardDuty,
		"inspector":   SecurityToolAwsRoleAdministratorFindingSecurityTools_Inspector,
		"macie":       SecurityToolAwsRoleAdministratorFindingSecurityTools_Macie,
		"securityhub": SecurityToolAwsRoleAdministratorFindingSecurityTools_SecurityHub,
		"wafshield":   SecurityToolAwsRoleAdministratorFindingSecurityTools_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityToolAwsRoleAdministratorFindingSecurityTools(input)
	return &out, nil
}
