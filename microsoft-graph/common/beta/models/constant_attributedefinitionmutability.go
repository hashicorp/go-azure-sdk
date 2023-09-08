package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionMutability string

const (
	AttributeDefinitionMutabilityImmutable AttributeDefinitionMutability = "Immutable"
	AttributeDefinitionMutabilityReadOnly  AttributeDefinitionMutability = "ReadOnly"
	AttributeDefinitionMutabilityReadWrite AttributeDefinitionMutability = "ReadWrite"
	AttributeDefinitionMutabilityWriteOnly AttributeDefinitionMutability = "WriteOnly"
)

func PossibleValuesForAttributeDefinitionMutability() []string {
	return []string{
		string(AttributeDefinitionMutabilityImmutable),
		string(AttributeDefinitionMutabilityReadOnly),
		string(AttributeDefinitionMutabilityReadWrite),
		string(AttributeDefinitionMutabilityWriteOnly),
	}
}

func parseAttributeDefinitionMutability(input string) (*AttributeDefinitionMutability, error) {
	vals := map[string]AttributeDefinitionMutability{
		"immutable": AttributeDefinitionMutabilityImmutable,
		"readonly":  AttributeDefinitionMutabilityReadOnly,
		"readwrite": AttributeDefinitionMutabilityReadWrite,
		"writeonly": AttributeDefinitionMutabilityWriteOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeDefinitionMutability(input)
	return &out, nil
}
