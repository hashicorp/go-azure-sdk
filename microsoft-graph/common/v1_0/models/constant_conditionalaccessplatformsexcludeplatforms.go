package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatformsExcludePlatforms string

const (
	ConditionalAccessPlatformsExcludePlatformsall          ConditionalAccessPlatformsExcludePlatforms = "All"
	ConditionalAccessPlatformsExcludePlatformsandroid      ConditionalAccessPlatformsExcludePlatforms = "Android"
	ConditionalAccessPlatformsExcludePlatformsiOS          ConditionalAccessPlatformsExcludePlatforms = "IOS"
	ConditionalAccessPlatformsExcludePlatformslinux        ConditionalAccessPlatformsExcludePlatforms = "Linux"
	ConditionalAccessPlatformsExcludePlatformsmacOS        ConditionalAccessPlatformsExcludePlatforms = "MacOS"
	ConditionalAccessPlatformsExcludePlatformswindows      ConditionalAccessPlatformsExcludePlatforms = "Windows"
	ConditionalAccessPlatformsExcludePlatformswindowsPhone ConditionalAccessPlatformsExcludePlatforms = "WindowsPhone"
)

func PossibleValuesForConditionalAccessPlatformsExcludePlatforms() []string {
	return []string{
		string(ConditionalAccessPlatformsExcludePlatformsall),
		string(ConditionalAccessPlatformsExcludePlatformsandroid),
		string(ConditionalAccessPlatformsExcludePlatformsiOS),
		string(ConditionalAccessPlatformsExcludePlatformslinux),
		string(ConditionalAccessPlatformsExcludePlatformsmacOS),
		string(ConditionalAccessPlatformsExcludePlatformswindows),
		string(ConditionalAccessPlatformsExcludePlatformswindowsPhone),
	}
}

func parseConditionalAccessPlatformsExcludePlatforms(input string) (*ConditionalAccessPlatformsExcludePlatforms, error) {
	vals := map[string]ConditionalAccessPlatformsExcludePlatforms{
		"all":          ConditionalAccessPlatformsExcludePlatformsall,
		"android":      ConditionalAccessPlatformsExcludePlatformsandroid,
		"ios":          ConditionalAccessPlatformsExcludePlatformsiOS,
		"linux":        ConditionalAccessPlatformsExcludePlatformslinux,
		"macos":        ConditionalAccessPlatformsExcludePlatformsmacOS,
		"windows":      ConditionalAccessPlatformsExcludePlatformswindows,
		"windowsphone": ConditionalAccessPlatformsExcludePlatformswindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPlatformsExcludePlatforms(input)
	return &out, nil
}
