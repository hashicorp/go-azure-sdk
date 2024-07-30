package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage string

const (
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Allowed       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "allowed"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Blocked       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "blocked"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_NotConfigured WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "notConfigured"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Required      WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "required"
)

func PossibleValuesForWindowsIdentityProtectionConfigurationPinSpecialCharactersUsage() []string {
	return []string{
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Allowed),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Blocked),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_NotConfigured),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Required),
	}
}

func (s *WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsIdentityProtectionConfigurationPinSpecialCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsIdentityProtectionConfigurationPinSpecialCharactersUsage(input string) (*WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage, error) {
	vals := map[string]WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage{
		"allowed":       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Allowed,
		"blocked":       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Blocked,
		"notconfigured": WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_NotConfigured,
		"required":      WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage(input)
	return &out, nil
}
