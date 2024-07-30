package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskApplicablePlatform string

const (
	SecurityConfigurationTaskApplicablePlatform_MacOS                     SecurityConfigurationTaskApplicablePlatform = "macOS"
	SecurityConfigurationTaskApplicablePlatform_Unknown                   SecurityConfigurationTaskApplicablePlatform = "unknown"
	SecurityConfigurationTaskApplicablePlatform_Windows10AndLater         SecurityConfigurationTaskApplicablePlatform = "windows10AndLater"
	SecurityConfigurationTaskApplicablePlatform_Windows10AndWindowsServer SecurityConfigurationTaskApplicablePlatform = "windows10AndWindowsServer"
)

func PossibleValuesForSecurityConfigurationTaskApplicablePlatform() []string {
	return []string{
		string(SecurityConfigurationTaskApplicablePlatform_MacOS),
		string(SecurityConfigurationTaskApplicablePlatform_Unknown),
		string(SecurityConfigurationTaskApplicablePlatform_Windows10AndLater),
		string(SecurityConfigurationTaskApplicablePlatform_Windows10AndWindowsServer),
	}
}

func (s *SecurityConfigurationTaskApplicablePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskApplicablePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskApplicablePlatform(input string) (*SecurityConfigurationTaskApplicablePlatform, error) {
	vals := map[string]SecurityConfigurationTaskApplicablePlatform{
		"macos":                     SecurityConfigurationTaskApplicablePlatform_MacOS,
		"unknown":                   SecurityConfigurationTaskApplicablePlatform_Unknown,
		"windows10andlater":         SecurityConfigurationTaskApplicablePlatform_Windows10AndLater,
		"windows10andwindowsserver": SecurityConfigurationTaskApplicablePlatform_Windows10AndWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskApplicablePlatform(input)
	return &out, nil
}
