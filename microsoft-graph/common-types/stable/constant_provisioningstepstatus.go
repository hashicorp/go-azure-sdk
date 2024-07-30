package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningStepStatus string

const (
	ProvisioningStepStatus_Failure ProvisioningStepStatus = "failure"
	ProvisioningStepStatus_Skipped ProvisioningStepStatus = "skipped"
	ProvisioningStepStatus_Success ProvisioningStepStatus = "success"
	ProvisioningStepStatus_Warning ProvisioningStepStatus = "warning"
)

func PossibleValuesForProvisioningStepStatus() []string {
	return []string{
		string(ProvisioningStepStatus_Failure),
		string(ProvisioningStepStatus_Skipped),
		string(ProvisioningStepStatus_Success),
		string(ProvisioningStepStatus_Warning),
	}
}

func (s *ProvisioningStepStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningStepStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningStepStatus(input string) (*ProvisioningStepStatus, error) {
	vals := map[string]ProvisioningStepStatus{
		"failure": ProvisioningStepStatus_Failure,
		"skipped": ProvisioningStepStatus_Skipped,
		"success": ProvisioningStepStatus_Success,
		"warning": ProvisioningStepStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStepStatus(input)
	return &out, nil
}
