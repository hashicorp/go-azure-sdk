package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess string

const (
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Disabled      WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "disabled"
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Enabled       WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "enabled"
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_NotConfigured WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "notConfigured"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Disabled),
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Enabled),
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_NotConfigured),
	}
}

func (s *WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess(input string) (*WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess{
		"disabled":      WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Disabled,
		"enabled":       WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_Enabled,
		"notconfigured": WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess(input)
	return &out, nil
}
