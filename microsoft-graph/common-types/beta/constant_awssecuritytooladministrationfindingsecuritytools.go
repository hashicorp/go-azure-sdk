package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecurityToolAdministrationFindingSecurityTools string

const (
	AwsSecurityToolAdministrationFindingSecurityTools_CloudTrail  AwsSecurityToolAdministrationFindingSecurityTools = "cloudTrail"
	AwsSecurityToolAdministrationFindingSecurityTools_Detective   AwsSecurityToolAdministrationFindingSecurityTools = "detective"
	AwsSecurityToolAdministrationFindingSecurityTools_GuardDuty   AwsSecurityToolAdministrationFindingSecurityTools = "guardDuty"
	AwsSecurityToolAdministrationFindingSecurityTools_Inspector   AwsSecurityToolAdministrationFindingSecurityTools = "inspector"
	AwsSecurityToolAdministrationFindingSecurityTools_Macie       AwsSecurityToolAdministrationFindingSecurityTools = "macie"
	AwsSecurityToolAdministrationFindingSecurityTools_SecurityHub AwsSecurityToolAdministrationFindingSecurityTools = "securityHub"
	AwsSecurityToolAdministrationFindingSecurityTools_WafShield   AwsSecurityToolAdministrationFindingSecurityTools = "wafShield"
)

func PossibleValuesForAwsSecurityToolAdministrationFindingSecurityTools() []string {
	return []string{
		string(AwsSecurityToolAdministrationFindingSecurityTools_CloudTrail),
		string(AwsSecurityToolAdministrationFindingSecurityTools_Detective),
		string(AwsSecurityToolAdministrationFindingSecurityTools_GuardDuty),
		string(AwsSecurityToolAdministrationFindingSecurityTools_Inspector),
		string(AwsSecurityToolAdministrationFindingSecurityTools_Macie),
		string(AwsSecurityToolAdministrationFindingSecurityTools_SecurityHub),
		string(AwsSecurityToolAdministrationFindingSecurityTools_WafShield),
	}
}

func (s *AwsSecurityToolAdministrationFindingSecurityTools) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsSecurityToolAdministrationFindingSecurityTools(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsSecurityToolAdministrationFindingSecurityTools(input string) (*AwsSecurityToolAdministrationFindingSecurityTools, error) {
	vals := map[string]AwsSecurityToolAdministrationFindingSecurityTools{
		"cloudtrail":  AwsSecurityToolAdministrationFindingSecurityTools_CloudTrail,
		"detective":   AwsSecurityToolAdministrationFindingSecurityTools_Detective,
		"guardduty":   AwsSecurityToolAdministrationFindingSecurityTools_GuardDuty,
		"inspector":   AwsSecurityToolAdministrationFindingSecurityTools_Inspector,
		"macie":       AwsSecurityToolAdministrationFindingSecurityTools_Macie,
		"securityhub": AwsSecurityToolAdministrationFindingSecurityTools_SecurityHub,
		"wafshield":   AwsSecurityToolAdministrationFindingSecurityTools_WafShield,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsSecurityToolAdministrationFindingSecurityTools(input)
	return &out, nil
}
