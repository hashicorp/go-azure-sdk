package deployments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentMode string

const (
	DeploymentModeComplete    DeploymentMode = "Complete"
	DeploymentModeIncremental DeploymentMode = "Incremental"
)

func PossibleValuesForDeploymentMode() []string {
	return []string{
		string(DeploymentModeComplete),
		string(DeploymentModeIncremental),
	}
}

func (s *DeploymentMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentMode(input string) (*DeploymentMode, error) {
	vals := map[string]DeploymentMode{
		"complete":    DeploymentModeComplete,
		"incremental": DeploymentModeIncremental,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentMode(input)
	return &out, nil
}
