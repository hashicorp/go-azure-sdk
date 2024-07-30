package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppInstallProgressDisplayLevel string

const (
	OfficeSuiteAppInstallProgressDisplayLevel_Full OfficeSuiteAppInstallProgressDisplayLevel = "full"
	OfficeSuiteAppInstallProgressDisplayLevel_None OfficeSuiteAppInstallProgressDisplayLevel = "none"
)

func PossibleValuesForOfficeSuiteAppInstallProgressDisplayLevel() []string {
	return []string{
		string(OfficeSuiteAppInstallProgressDisplayLevel_Full),
		string(OfficeSuiteAppInstallProgressDisplayLevel_None),
	}
}

func (s *OfficeSuiteAppInstallProgressDisplayLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeSuiteAppInstallProgressDisplayLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeSuiteAppInstallProgressDisplayLevel(input string) (*OfficeSuiteAppInstallProgressDisplayLevel, error) {
	vals := map[string]OfficeSuiteAppInstallProgressDisplayLevel{
		"full": OfficeSuiteAppInstallProgressDisplayLevel_Full,
		"none": OfficeSuiteAppInstallProgressDisplayLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppInstallProgressDisplayLevel(input)
	return &out, nil
}
