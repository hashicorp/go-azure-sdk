package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipOperationType string

const (
	DelegatedAdminRelationshipOperationType_DelegatedAdminAccessAssignmentUpdate DelegatedAdminRelationshipOperationType = "delegatedAdminAccessAssignmentUpdate"
)

func PossibleValuesForDelegatedAdminRelationshipOperationType() []string {
	return []string{
		string(DelegatedAdminRelationshipOperationType_DelegatedAdminAccessAssignmentUpdate),
	}
}

func (s *DelegatedAdminRelationshipOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipOperationType(input string) (*DelegatedAdminRelationshipOperationType, error) {
	vals := map[string]DelegatedAdminRelationshipOperationType{
		"delegatedadminaccessassignmentupdate": DelegatedAdminRelationshipOperationType_DelegatedAdminAccessAssignmentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipOperationType(input)
	return &out, nil
}
