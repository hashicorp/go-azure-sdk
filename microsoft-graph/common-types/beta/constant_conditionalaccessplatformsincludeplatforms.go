package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatformsIncludePlatforms string

const (
	ConditionalAccessPlatformsIncludePlatforms_All          ConditionalAccessPlatformsIncludePlatforms = "all"
	ConditionalAccessPlatformsIncludePlatforms_Android      ConditionalAccessPlatformsIncludePlatforms = "android"
	ConditionalAccessPlatformsIncludePlatforms_IOS          ConditionalAccessPlatformsIncludePlatforms = "iOS"
	ConditionalAccessPlatformsIncludePlatforms_Linux        ConditionalAccessPlatformsIncludePlatforms = "linux"
	ConditionalAccessPlatformsIncludePlatforms_MacOS        ConditionalAccessPlatformsIncludePlatforms = "macOS"
	ConditionalAccessPlatformsIncludePlatforms_Windows      ConditionalAccessPlatformsIncludePlatforms = "windows"
	ConditionalAccessPlatformsIncludePlatforms_WindowsPhone ConditionalAccessPlatformsIncludePlatforms = "windowsPhone"
)

func PossibleValuesForConditionalAccessPlatformsIncludePlatforms() []string {
	return []string{
		string(ConditionalAccessPlatformsIncludePlatforms_All),
		string(ConditionalAccessPlatformsIncludePlatforms_Android),
		string(ConditionalAccessPlatformsIncludePlatforms_IOS),
		string(ConditionalAccessPlatformsIncludePlatforms_Linux),
		string(ConditionalAccessPlatformsIncludePlatforms_MacOS),
		string(ConditionalAccessPlatformsIncludePlatforms_Windows),
		string(ConditionalAccessPlatformsIncludePlatforms_WindowsPhone),
	}
}

func (s *ConditionalAccessPlatformsIncludePlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessPlatformsIncludePlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessPlatformsIncludePlatforms(input string) (*ConditionalAccessPlatformsIncludePlatforms, error) {
	vals := map[string]ConditionalAccessPlatformsIncludePlatforms{
		"all":          ConditionalAccessPlatformsIncludePlatforms_All,
		"android":      ConditionalAccessPlatformsIncludePlatforms_Android,
		"ios":          ConditionalAccessPlatformsIncludePlatforms_IOS,
		"linux":        ConditionalAccessPlatformsIncludePlatforms_Linux,
		"macos":        ConditionalAccessPlatformsIncludePlatforms_MacOS,
		"windows":      ConditionalAccessPlatformsIncludePlatforms_Windows,
		"windowsphone": ConditionalAccessPlatformsIncludePlatforms_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPlatformsIncludePlatforms(input)
	return &out, nil
}
