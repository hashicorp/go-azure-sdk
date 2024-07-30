package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityToolAwsResourceAdministratorFindingSecurityTools string

const (
	SecurityToolAwsResourceAdministratorFindingSecurityTools_CloudTrail  SecurityToolAwsResourceAdministratorFindingSecurityTools = "cloudTrail"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_Detective   SecurityToolAwsResourceAdministratorFindingSecurityTools = "detective"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_GuardDuty   SecurityToolAwsResourceAdministratorFindingSecurityTools = "guardDuty"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_Inspector   SecurityToolAwsResourceAdministratorFindingSecurityTools = "inspector"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_Macie       SecurityToolAwsResourceAdministratorFindingSecurityTools = "macie"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_SecurityHub SecurityToolAwsResourceAdministratorFindingSecurityTools = "securityHub"
	SecurityToolAwsResourceAdministratorFindingSecurityTools_WafShield   SecurityToolAwsResourceAdministratorFindingSecurityTools = "wafShield"
)

func PossibleValuesForSecurityToolAwsResourceAdministratorFindingSecurityTools() []string {
	return []string{
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_CloudTrail),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_Detective),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_GuardDuty),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_Inspector),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_Macie),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_SecurityHub),
		string(SecurityToolAwsResourceAdministratorFindingSecurityTools_WafShield),
	}
}

func (s *SecurityToolAwsResourceAdministratorFindingSecurityTools) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityToolAwsResourceAdministratorFindingSecurityTools(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityToolAwsResourceAdministratorFindingSecurityTools(input string) (*SecurityToolAwsResourceAdministratorFindingSecurityTools, error) {
	vals := map[string]SecurityToolAwsResourceAdministratorFindingSecurityTools{
		"cloudtrail":  SecurityToolAwsResourceAdministratorFindingSecurityTools_CloudTrail,
		"detective":   SecurityToolAwsResourceAdministratorFindingSecurityTools_Detective,
		"guardduty":   SecurityToolAwsResourceAdministratorFindingSecurityTools_GuardDuty,
		"inspector":   SecurityToolAwsResourceAdministratorFindingSecurityTools_Inspector,
		"macie":       SecurityToolAwsResourceAdministratorFindingSecurityTools_Macie,
		"securityhub": SecurityToolAwsResourceAdministratorFindingSecurityTools_SecurityHub,
		"wafshield":   SecurityToolAwsResourceAdministratorFindingSecurityTools_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityToolAwsResourceAdministratorFindingSecurityTools(input)
	return &out, nil
}
