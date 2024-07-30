package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationInkWorkspaceAccessState string

const (
	Windows10GeneralConfigurationInkWorkspaceAccessState_Allowed       Windows10GeneralConfigurationInkWorkspaceAccessState = "allowed"
	Windows10GeneralConfigurationInkWorkspaceAccessState_Blocked       Windows10GeneralConfigurationInkWorkspaceAccessState = "blocked"
	Windows10GeneralConfigurationInkWorkspaceAccessState_NotConfigured Windows10GeneralConfigurationInkWorkspaceAccessState = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationInkWorkspaceAccessState() []string {
	return []string{
		string(Windows10GeneralConfigurationInkWorkspaceAccessState_Allowed),
		string(Windows10GeneralConfigurationInkWorkspaceAccessState_Blocked),
		string(Windows10GeneralConfigurationInkWorkspaceAccessState_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationInkWorkspaceAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationInkWorkspaceAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationInkWorkspaceAccessState(input string) (*Windows10GeneralConfigurationInkWorkspaceAccessState, error) {
	vals := map[string]Windows10GeneralConfigurationInkWorkspaceAccessState{
		"allowed":       Windows10GeneralConfigurationInkWorkspaceAccessState_Allowed,
		"blocked":       Windows10GeneralConfigurationInkWorkspaceAccessState_Blocked,
		"notconfigured": Windows10GeneralConfigurationInkWorkspaceAccessState_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationInkWorkspaceAccessState(input)
	return &out, nil
}
