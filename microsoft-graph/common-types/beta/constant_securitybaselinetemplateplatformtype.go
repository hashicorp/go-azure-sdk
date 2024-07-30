package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplatePlatformType string

const (
	SecurityBaselineTemplatePlatformType_All                SecurityBaselineTemplatePlatformType = "all"
	SecurityBaselineTemplatePlatformType_Android            SecurityBaselineTemplatePlatformType = "android"
	SecurityBaselineTemplatePlatformType_AndroidAOSP        SecurityBaselineTemplatePlatformType = "androidAOSP"
	SecurityBaselineTemplatePlatformType_AndroidForWork     SecurityBaselineTemplatePlatformType = "androidForWork"
	SecurityBaselineTemplatePlatformType_AndroidWorkProfile SecurityBaselineTemplatePlatformType = "androidWorkProfile"
	SecurityBaselineTemplatePlatformType_IOS                SecurityBaselineTemplatePlatformType = "iOS"
	SecurityBaselineTemplatePlatformType_MacOS              SecurityBaselineTemplatePlatformType = "macOS"
	SecurityBaselineTemplatePlatformType_Windows10AndLater  SecurityBaselineTemplatePlatformType = "windows10AndLater"
	SecurityBaselineTemplatePlatformType_Windows10XProfile  SecurityBaselineTemplatePlatformType = "windows10XProfile"
	SecurityBaselineTemplatePlatformType_Windows81AndLater  SecurityBaselineTemplatePlatformType = "windows81AndLater"
	SecurityBaselineTemplatePlatformType_WindowsPhone81     SecurityBaselineTemplatePlatformType = "windowsPhone81"
)

func PossibleValuesForSecurityBaselineTemplatePlatformType() []string {
	return []string{
		string(SecurityBaselineTemplatePlatformType_All),
		string(SecurityBaselineTemplatePlatformType_Android),
		string(SecurityBaselineTemplatePlatformType_AndroidAOSP),
		string(SecurityBaselineTemplatePlatformType_AndroidForWork),
		string(SecurityBaselineTemplatePlatformType_AndroidWorkProfile),
		string(SecurityBaselineTemplatePlatformType_IOS),
		string(SecurityBaselineTemplatePlatformType_MacOS),
		string(SecurityBaselineTemplatePlatformType_Windows10AndLater),
		string(SecurityBaselineTemplatePlatformType_Windows10XProfile),
		string(SecurityBaselineTemplatePlatformType_Windows81AndLater),
		string(SecurityBaselineTemplatePlatformType_WindowsPhone81),
	}
}

func (s *SecurityBaselineTemplatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineTemplatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineTemplatePlatformType(input string) (*SecurityBaselineTemplatePlatformType, error) {
	vals := map[string]SecurityBaselineTemplatePlatformType{
		"all":                SecurityBaselineTemplatePlatformType_All,
		"android":            SecurityBaselineTemplatePlatformType_Android,
		"androidaosp":        SecurityBaselineTemplatePlatformType_AndroidAOSP,
		"androidforwork":     SecurityBaselineTemplatePlatformType_AndroidForWork,
		"androidworkprofile": SecurityBaselineTemplatePlatformType_AndroidWorkProfile,
		"ios":                SecurityBaselineTemplatePlatformType_IOS,
		"macos":              SecurityBaselineTemplatePlatformType_MacOS,
		"windows10andlater":  SecurityBaselineTemplatePlatformType_Windows10AndLater,
		"windows10xprofile":  SecurityBaselineTemplatePlatformType_Windows10XProfile,
		"windows81andlater":  SecurityBaselineTemplatePlatformType_Windows81AndLater,
		"windowsphone81":     SecurityBaselineTemplatePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplatePlatformType(input)
	return &out, nil
}
