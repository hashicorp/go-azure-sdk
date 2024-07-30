package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperationOperationType string

const (
	DelegatedAdminRelationshipOperationOperationType_DelegatedAdminAccessAssignmentUpdate DelegatedAdminRelationshipOperationOperationType = "delegatedAdminAccessAssignmentUpdate"
)

func PossibleValuesForDelegatedAdminRelationshipOperationOperationType() []string {
	return []string{
		string(DelegatedAdminRelationshipOperationOperationType_DelegatedAdminAccessAssignmentUpdate),
	}
}

func (s *DelegatedAdminRelationshipOperationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipOperationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipOperationOperationType(input string) (*DelegatedAdminRelationshipOperationOperationType, error) {
	vals := map[string]DelegatedAdminRelationshipOperationOperationType{
		"delegatedadminaccessassignmentupdate": DelegatedAdminRelationshipOperationOperationType_DelegatedAdminAccessAssignmentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipOperationOperationType(input)
	return &out, nil
}
