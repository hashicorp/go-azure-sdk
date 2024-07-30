package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPlatformsExcludePlatforms string

const (
	ConditionalAccessPlatformsExcludePlatforms_All          ConditionalAccessPlatformsExcludePlatforms = "all"
	ConditionalAccessPlatformsExcludePlatforms_Android      ConditionalAccessPlatformsExcludePlatforms = "android"
	ConditionalAccessPlatformsExcludePlatforms_IOS          ConditionalAccessPlatformsExcludePlatforms = "iOS"
	ConditionalAccessPlatformsExcludePlatforms_Linux        ConditionalAccessPlatformsExcludePlatforms = "linux"
	ConditionalAccessPlatformsExcludePlatforms_MacOS        ConditionalAccessPlatformsExcludePlatforms = "macOS"
	ConditionalAccessPlatformsExcludePlatforms_Windows      ConditionalAccessPlatformsExcludePlatforms = "windows"
	ConditionalAccessPlatformsExcludePlatforms_WindowsPhone ConditionalAccessPlatformsExcludePlatforms = "windowsPhone"
)

func PossibleValuesForConditionalAccessPlatformsExcludePlatforms() []string {
	return []string{
		string(ConditionalAccessPlatformsExcludePlatforms_All),
		string(ConditionalAccessPlatformsExcludePlatforms_Android),
		string(ConditionalAccessPlatformsExcludePlatforms_IOS),
		string(ConditionalAccessPlatformsExcludePlatforms_Linux),
		string(ConditionalAccessPlatformsExcludePlatforms_MacOS),
		string(ConditionalAccessPlatformsExcludePlatforms_Windows),
		string(ConditionalAccessPlatformsExcludePlatforms_WindowsPhone),
	}
}

func (s *ConditionalAccessPlatformsExcludePlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessPlatformsExcludePlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessPlatformsExcludePlatforms(input string) (*ConditionalAccessPlatformsExcludePlatforms, error) {
	vals := map[string]ConditionalAccessPlatformsExcludePlatforms{
		"all":          ConditionalAccessPlatformsExcludePlatforms_All,
		"android":      ConditionalAccessPlatformsExcludePlatforms_Android,
		"ios":          ConditionalAccessPlatformsExcludePlatforms_IOS,
		"linux":        ConditionalAccessPlatformsExcludePlatforms_Linux,
		"macos":        ConditionalAccessPlatformsExcludePlatforms_MacOS,
		"windows":      ConditionalAccessPlatformsExcludePlatforms_Windows,
		"windowsphone": ConditionalAccessPlatformsExcludePlatforms_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPlatformsExcludePlatforms(input)
	return &out, nil
}
