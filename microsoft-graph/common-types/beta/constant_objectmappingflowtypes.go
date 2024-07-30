package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingFlowTypes string

const (
	ObjectMappingFlowTypes_Add    ObjectMappingFlowTypes = "Add"
	ObjectMappingFlowTypes_Delete ObjectMappingFlowTypes = "Delete"
	ObjectMappingFlowTypes_None   ObjectMappingFlowTypes = "None"
	ObjectMappingFlowTypes_Update ObjectMappingFlowTypes = "Update"
)

func PossibleValuesForObjectMappingFlowTypes() []string {
	return []string{
		string(ObjectMappingFlowTypes_Add),
		string(ObjectMappingFlowTypes_Delete),
		string(ObjectMappingFlowTypes_None),
		string(ObjectMappingFlowTypes_Update),
	}
}

func (s *ObjectMappingFlowTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObjectMappingFlowTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObjectMappingFlowTypes(input string) (*ObjectMappingFlowTypes, error) {
	vals := map[string]ObjectMappingFlowTypes{
		"add":    ObjectMappingFlowTypes_Add,
		"delete": ObjectMappingFlowTypes_Delete,
		"none":   ObjectMappingFlowTypes_None,
		"update": ObjectMappingFlowTypes_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingFlowTypes(input)
	return &out, nil
}
