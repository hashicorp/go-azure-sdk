package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMutability string

const (
	AttributeDefinitionMutability_Immutable AttributeDefinitionMutability = "Immutable"
	AttributeDefinitionMutability_ReadOnly  AttributeDefinitionMutability = "ReadOnly"
	AttributeDefinitionMutability_ReadWrite AttributeDefinitionMutability = "ReadWrite"
	AttributeDefinitionMutability_WriteOnly AttributeDefinitionMutability = "WriteOnly"
)

func PossibleValuesForAttributeDefinitionMutability() []string {
	return []string{
		string(AttributeDefinitionMutability_Immutable),
		string(AttributeDefinitionMutability_ReadOnly),
		string(AttributeDefinitionMutability_ReadWrite),
		string(AttributeDefinitionMutability_WriteOnly),
	}
}

func (s *AttributeDefinitionMutability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeDefinitionMutability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeDefinitionMutability(input string) (*AttributeDefinitionMutability, error) {
	vals := map[string]AttributeDefinitionMutability{
		"immutable": AttributeDefinitionMutability_Immutable,
		"readonly":  AttributeDefinitionMutability_ReadOnly,
		"readwrite": AttributeDefinitionMutability_ReadWrite,
		"writeonly": AttributeDefinitionMutability_WriteOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMutability(input)
	return &out, nil
}
