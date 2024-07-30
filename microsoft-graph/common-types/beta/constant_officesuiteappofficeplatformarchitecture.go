package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppOfficePlatformArchitecture string

const (
	OfficeSuiteAppOfficePlatformArchitecture_Arm     OfficeSuiteAppOfficePlatformArchitecture = "arm"
	OfficeSuiteAppOfficePlatformArchitecture_Arm64   OfficeSuiteAppOfficePlatformArchitecture = "arm64"
	OfficeSuiteAppOfficePlatformArchitecture_Neutral OfficeSuiteAppOfficePlatformArchitecture = "neutral"
	OfficeSuiteAppOfficePlatformArchitecture_None    OfficeSuiteAppOfficePlatformArchitecture = "none"
	OfficeSuiteAppOfficePlatformArchitecture_X64     OfficeSuiteAppOfficePlatformArchitecture = "x64"
	OfficeSuiteAppOfficePlatformArchitecture_X86     OfficeSuiteAppOfficePlatformArchitecture = "x86"
)

func PossibleValuesForOfficeSuiteAppOfficePlatformArchitecture() []string {
	return []string{
		string(OfficeSuiteAppOfficePlatformArchitecture_Arm),
		string(OfficeSuiteAppOfficePlatformArchitecture_Arm64),
		string(OfficeSuiteAppOfficePlatformArchitecture_Neutral),
		string(OfficeSuiteAppOfficePlatformArchitecture_None),
		string(OfficeSuiteAppOfficePlatformArchitecture_X64),
		string(OfficeSuiteAppOfficePlatformArchitecture_X86),
	}
}

func (s *OfficeSuiteAppOfficePlatformArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppOfficePlatformArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppOfficePlatformArchitecture(input string) (*OfficeSuiteAppOfficePlatformArchitecture, error) {
	vals := map[string]OfficeSuiteAppOfficePlatformArchitecture{
		"arm":     OfficeSuiteAppOfficePlatformArchitecture_Arm,
		"arm64":   OfficeSuiteAppOfficePlatformArchitecture_Arm64,
		"neutral": OfficeSuiteAppOfficePlatformArchitecture_Neutral,
		"none":    OfficeSuiteAppOfficePlatformArchitecture_None,
		"x64":     OfficeSuiteAppOfficePlatformArchitecture_X64,
		"x86":     OfficeSuiteAppOfficePlatformArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppOfficePlatformArchitecture(input)
	return &out, nil
}
