package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingTokenType string

const (
	DepOnboardingSettingTokenTypeappleSchoolManager DepOnboardingSettingTokenType = "AppleSchoolManager"
	DepOnboardingSettingTokenTypedep                DepOnboardingSettingTokenType = "Dep"
	DepOnboardingSettingTokenTypenone               DepOnboardingSettingTokenType = "None"
)

func PossibleValuesForDepOnboardingSettingTokenType() []string {
	return []string{
		string(DepOnboardingSettingTokenTypeappleSchoolManager),
		string(DepOnboardingSettingTokenTypedep),
		string(DepOnboardingSettingTokenTypenone),
	}
}

func parseDepOnboardingSettingTokenType(input string) (*DepOnboardingSettingTokenType, error) {
	vals := map[string]DepOnboardingSettingTokenType{
		"appleschoolmanager": DepOnboardingSettingTokenTypeappleSchoolManager,
		"dep":                DepOnboardingSettingTokenTypedep,
		"none":               DepOnboardingSettingTokenTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepOnboardingSettingTokenType(input)
	return &out, nil
}
