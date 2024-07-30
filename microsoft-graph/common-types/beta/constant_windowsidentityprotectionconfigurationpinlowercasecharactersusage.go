package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage string

const (
	WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Allowed       WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage = "allowed"
	WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Blocked       WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage = "blocked"
	WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_NotConfigured WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage = "notConfigured"
	WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Required      WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage = "required"
)

func PossibleValuesForWindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage() []string {
	return []string{
		string(WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Allowed),
		string(WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Blocked),
		string(WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_NotConfigured),
		string(WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Required),
	}
}

func (s *WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage(input string) (*WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage, error) {
	vals := map[string]WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage{
		"allowed":       WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Allowed,
		"blocked":       WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Blocked,
		"notconfigured": WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_NotConfigured,
		"required":      WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage(input)
	return &out, nil
}
