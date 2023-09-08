package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettingsImageTaggingOption string

const (
	SharepointSettingsImageTaggingOptionbasic    SharepointSettingsImageTaggingOption = "Basic"
	SharepointSettingsImageTaggingOptiondisabled SharepointSettingsImageTaggingOption = "Disabled"
	SharepointSettingsImageTaggingOptionenhanced SharepointSettingsImageTaggingOption = "Enhanced"
)

func PossibleValuesForSharepointSettingsImageTaggingOption() []string {
	return []string{
		string(SharepointSettingsImageTaggingOptionbasic),
		string(SharepointSettingsImageTaggingOptiondisabled),
		string(SharepointSettingsImageTaggingOptionenhanced),
	}
}

func parseSharepointSettingsImageTaggingOption(input string) (*SharepointSettingsImageTaggingOption, error) {
	vals := map[string]SharepointSettingsImageTaggingOption{
		"basic":    SharepointSettingsImageTaggingOptionbasic,
		"disabled": SharepointSettingsImageTaggingOptiondisabled,
		"enhanced": SharepointSettingsImageTaggingOptionenhanced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharepointSettingsImageTaggingOption(input)
	return &out, nil
}
