package lifecycleworkflowdeleteditemworkflowruntaskprocessingresult

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus string

const (
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusCompleted IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "completed"
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusFailed    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "failed"
)

func PossibleValuesForIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus() []string {
	return []string{
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusCompleted),
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusFailed),
	}
}

func (s *IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(input string) (*IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus, error) {
	vals := map[string]IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus{
		"completed": IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusCompleted,
		"failed":    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(input)
	return &out, nil
}
