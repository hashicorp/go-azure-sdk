package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppManagedInstaller string

const (
	WindowsManagementAppManagedInstaller_Disabled WindowsManagementAppManagedInstaller = "disabled"
	WindowsManagementAppManagedInstaller_Enabled  WindowsManagementAppManagedInstaller = "enabled"
)

func PossibleValuesForWindowsManagementAppManagedInstaller() []string {
	return []string{
		string(WindowsManagementAppManagedInstaller_Disabled),
		string(WindowsManagementAppManagedInstaller_Enabled),
	}
}

func (s *WindowsManagementAppManagedInstaller) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagementAppManagedInstaller(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagementAppManagedInstaller(input string) (*WindowsManagementAppManagedInstaller, error) {
	vals := map[string]WindowsManagementAppManagedInstaller{
		"disabled": WindowsManagementAppManagedInstaller_Disabled,
		"enabled":  WindowsManagementAppManagedInstaller_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagementAppManagedInstaller(input)
	return &out, nil
}
