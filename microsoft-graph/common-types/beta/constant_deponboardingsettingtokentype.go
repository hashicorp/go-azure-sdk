package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingTokenType string

const (
	DepOnboardingSettingTokenType_AppleSchoolManager DepOnboardingSettingTokenType = "appleSchoolManager"
	DepOnboardingSettingTokenType_Dep                DepOnboardingSettingTokenType = "dep"
	DepOnboardingSettingTokenType_None               DepOnboardingSettingTokenType = "none"
)

func PossibleValuesForDepOnboardingSettingTokenType() []string {
	return []string{
		string(DepOnboardingSettingTokenType_AppleSchoolManager),
		string(DepOnboardingSettingTokenType_Dep),
		string(DepOnboardingSettingTokenType_None),
	}
}

func (s *DepOnboardingSettingTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDepOnboardingSettingTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDepOnboardingSettingTokenType(input string) (*DepOnboardingSettingTokenType, error) {
	vals := map[string]DepOnboardingSettingTokenType{
		"appleschoolmanager": DepOnboardingSettingTokenType_AppleSchoolManager,
		"dep":                DepOnboardingSettingTokenType_Dep,
		"none":               DepOnboardingSettingTokenType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepOnboardingSettingTokenType(input)
	return &out, nil
}
