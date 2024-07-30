package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationInkWorkspaceAccess string

const (
	Windows10GeneralConfigurationInkWorkspaceAccess_Disabled      Windows10GeneralConfigurationInkWorkspaceAccess = "disabled"
	Windows10GeneralConfigurationInkWorkspaceAccess_Enabled       Windows10GeneralConfigurationInkWorkspaceAccess = "enabled"
	Windows10GeneralConfigurationInkWorkspaceAccess_NotConfigured Windows10GeneralConfigurationInkWorkspaceAccess = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationInkWorkspaceAccess() []string {
	return []string{
		string(Windows10GeneralConfigurationInkWorkspaceAccess_Disabled),
		string(Windows10GeneralConfigurationInkWorkspaceAccess_Enabled),
		string(Windows10GeneralConfigurationInkWorkspaceAccess_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationInkWorkspaceAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationInkWorkspaceAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationInkWorkspaceAccess(input string) (*Windows10GeneralConfigurationInkWorkspaceAccess, error) {
	vals := map[string]Windows10GeneralConfigurationInkWorkspaceAccess{
		"disabled":      Windows10GeneralConfigurationInkWorkspaceAccess_Disabled,
		"enabled":       Windows10GeneralConfigurationInkWorkspaceAccess_Enabled,
		"notconfigured": Windows10GeneralConfigurationInkWorkspaceAccess_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationInkWorkspaceAccess(input)
	return &out, nil
}
