package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingFlowType string

const (
	AttributeMappingFlowType_Always            AttributeMappingFlowType = "Always"
	AttributeMappingFlowType_AttributeAddOnly  AttributeMappingFlowType = "AttributeAddOnly"
	AttributeMappingFlowType_MultiValueAddOnly AttributeMappingFlowType = "MultiValueAddOnly"
	AttributeMappingFlowType_ObjectAddOnly     AttributeMappingFlowType = "ObjectAddOnly"
	AttributeMappingFlowType_ValueAddOnly      AttributeMappingFlowType = "ValueAddOnly"
)

func PossibleValuesForAttributeMappingFlowType() []string {
	return []string{
		string(AttributeMappingFlowType_Always),
		string(AttributeMappingFlowType_AttributeAddOnly),
		string(AttributeMappingFlowType_MultiValueAddOnly),
		string(AttributeMappingFlowType_ObjectAddOnly),
		string(AttributeMappingFlowType_ValueAddOnly),
	}
}

func (s *AttributeMappingFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingFlowType(input string) (*AttributeMappingFlowType, error) {
	vals := map[string]AttributeMappingFlowType{
		"always":            AttributeMappingFlowType_Always,
		"attributeaddonly":  AttributeMappingFlowType_AttributeAddOnly,
		"multivalueaddonly": AttributeMappingFlowType_MultiValueAddOnly,
		"objectaddonly":     AttributeMappingFlowType_ObjectAddOnly,
		"valueaddonly":      AttributeMappingFlowType_ValueAddOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowType(input)
	return &out, nil
}
