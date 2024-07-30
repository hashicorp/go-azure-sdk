package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentProfileAssignmentSource string

const (
	WindowsAutopilotDeploymentProfileAssignmentSource_Direct     WindowsAutopilotDeploymentProfileAssignmentSource = "direct"
	WindowsAutopilotDeploymentProfileAssignmentSource_PolicySets WindowsAutopilotDeploymentProfileAssignmentSource = "policySets"
)

func PossibleValuesForWindowsAutopilotDeploymentProfileAssignmentSource() []string {
	return []string{
		string(WindowsAutopilotDeploymentProfileAssignmentSource_Direct),
		string(WindowsAutopilotDeploymentProfileAssignmentSource_PolicySets),
	}
}

func (s *WindowsAutopilotDeploymentProfileAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeploymentProfileAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeploymentProfileAssignmentSource(input string) (*WindowsAutopilotDeploymentProfileAssignmentSource, error) {
	vals := map[string]WindowsAutopilotDeploymentProfileAssignmentSource{
		"direct":     WindowsAutopilotDeploymentProfileAssignmentSource_Direct,
		"policysets": WindowsAutopilotDeploymentProfileAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentProfileAssignmentSource(input)
	return &out, nil
}
