package scriptpackages

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptPackageProvisioningState string

const (
	ScriptPackageProvisioningStateCanceled  ScriptPackageProvisioningState = "Canceled"
	ScriptPackageProvisioningStateFailed    ScriptPackageProvisioningState = "Failed"
	ScriptPackageProvisioningStateSucceeded ScriptPackageProvisioningState = "Succeeded"
)

func PossibleValuesForScriptPackageProvisioningState() []string {
	return []string{
		string(ScriptPackageProvisioningStateCanceled),
		string(ScriptPackageProvisioningStateFailed),
		string(ScriptPackageProvisioningStateSucceeded),
	}
}

func (s *ScriptPackageProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScriptPackageProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScriptPackageProvisioningState(input string) (*ScriptPackageProvisioningState, error) {
	vals := map[string]ScriptPackageProvisioningState{
		"canceled":  ScriptPackageProvisioningStateCanceled,
		"failed":    ScriptPackageProvisioningStateFailed,
		"succeeded": ScriptPackageProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScriptPackageProvisioningState(input)
	return &out, nil
}
