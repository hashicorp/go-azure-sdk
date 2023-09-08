package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatformsIncludePlatforms string

const (
	ConditionalAccessPlatformsIncludePlatformsall          ConditionalAccessPlatformsIncludePlatforms = "All"
	ConditionalAccessPlatformsIncludePlatformsandroid      ConditionalAccessPlatformsIncludePlatforms = "Android"
	ConditionalAccessPlatformsIncludePlatformsiOS          ConditionalAccessPlatformsIncludePlatforms = "IOS"
	ConditionalAccessPlatformsIncludePlatformslinux        ConditionalAccessPlatformsIncludePlatforms = "Linux"
	ConditionalAccessPlatformsIncludePlatformsmacOS        ConditionalAccessPlatformsIncludePlatforms = "MacOS"
	ConditionalAccessPlatformsIncludePlatformswindows      ConditionalAccessPlatformsIncludePlatforms = "Windows"
	ConditionalAccessPlatformsIncludePlatformswindowsPhone ConditionalAccessPlatformsIncludePlatforms = "WindowsPhone"
)

func PossibleValuesForConditionalAccessPlatformsIncludePlatforms() []string {
	return []string{
		string(ConditionalAccessPlatformsIncludePlatformsall),
		string(ConditionalAccessPlatformsIncludePlatformsandroid),
		string(ConditionalAccessPlatformsIncludePlatformsiOS),
		string(ConditionalAccessPlatformsIncludePlatformslinux),
		string(ConditionalAccessPlatformsIncludePlatformsmacOS),
		string(ConditionalAccessPlatformsIncludePlatformswindows),
		string(ConditionalAccessPlatformsIncludePlatformswindowsPhone),
	}
}

func parseConditionalAccessPlatformsIncludePlatforms(input string) (*ConditionalAccessPlatformsIncludePlatforms, error) {
	vals := map[string]ConditionalAccessPlatformsIncludePlatforms{
		"all":          ConditionalAccessPlatformsIncludePlatformsall,
		"android":      ConditionalAccessPlatformsIncludePlatformsandroid,
		"ios":          ConditionalAccessPlatformsIncludePlatformsiOS,
		"linux":        ConditionalAccessPlatformsIncludePlatformslinux,
		"macos":        ConditionalAccessPlatformsIncludePlatformsmacOS,
		"windows":      ConditionalAccessPlatformsIncludePlatformswindows,
		"windowsphone": ConditionalAccessPlatformsIncludePlatformswindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPlatformsIncludePlatforms(input)
	return &out, nil
}
