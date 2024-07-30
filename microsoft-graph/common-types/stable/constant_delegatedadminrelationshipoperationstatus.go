package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperationStatus string

const (
	DelegatedAdminRelationshipOperationStatus_Failed     DelegatedAdminRelationshipOperationStatus = "failed"
	DelegatedAdminRelationshipOperationStatus_NotStarted DelegatedAdminRelationshipOperationStatus = "notStarted"
	DelegatedAdminRelationshipOperationStatus_Running    DelegatedAdminRelationshipOperationStatus = "running"
	DelegatedAdminRelationshipOperationStatus_Succeeded  DelegatedAdminRelationshipOperationStatus = "succeeded"
)

func PossibleValuesForDelegatedAdminRelationshipOperationStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipOperationStatus_Failed),
		string(DelegatedAdminRelationshipOperationStatus_NotStarted),
		string(DelegatedAdminRelationshipOperationStatus_Running),
		string(DelegatedAdminRelationshipOperationStatus_Succeeded),
	}
}

func (s *DelegatedAdminRelationshipOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipOperationStatus(input string) (*DelegatedAdminRelationshipOperationStatus, error) {
	vals := map[string]DelegatedAdminRelationshipOperationStatus{
		"failed":     DelegatedAdminRelationshipOperationStatus_Failed,
		"notstarted": DelegatedAdminRelationshipOperationStatus_NotStarted,
		"running":    DelegatedAdminRelationshipOperationStatus_Running,
		"succeeded":  DelegatedAdminRelationshipOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipOperationStatus(input)
	return &out, nil
}
