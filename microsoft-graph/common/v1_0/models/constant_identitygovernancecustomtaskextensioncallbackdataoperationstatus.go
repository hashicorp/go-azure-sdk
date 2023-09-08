package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus string

const (
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatuscompleted IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "Completed"
	IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusfailed    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus = "Failed"
)

func PossibleValuesForIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus() []string {
	return []string{
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatuscompleted),
		string(IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusfailed),
	}
}

func parseIdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(input string) (*IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus, error) {
	vals := map[string]IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus{
		"completed": IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatuscompleted,
		"failed":    IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatusfailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus(input)
	return &out, nil
}
