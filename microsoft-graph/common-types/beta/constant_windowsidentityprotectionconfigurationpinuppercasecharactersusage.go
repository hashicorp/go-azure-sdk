package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage string

const (
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Allowed       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "allowed"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Blocked       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "blocked"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_NotConfigured WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "notConfigured"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Required      WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "required"
)

func PossibleValuesForWindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage() []string {
	return []string{
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Allowed),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Blocked),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_NotConfigured),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Required),
	}
}

func (s *WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage(input string) (*WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage, error) {
	vals := map[string]WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage{
		"allowed":       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Allowed,
		"blocked":       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Blocked,
		"notconfigured": WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_NotConfigured,
		"required":      WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage(input)
	return &out, nil
}
