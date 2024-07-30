package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus string

const (
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Completed IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "completed"
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Failed    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "failed"
)

func PossibleValuesForIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus() []string {
	return []string{
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Completed),
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Failed),
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
		"completed": IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Completed,
		"failed":    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus_Failed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(input)
	return &out, nil
}
