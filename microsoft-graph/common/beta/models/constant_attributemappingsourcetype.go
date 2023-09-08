package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingSourceType string

const (
	AttributeMappingSourceTypeAttribute AttributeMappingSourceType = "Attribute"
	AttributeMappingSourceTypeConstant  AttributeMappingSourceType = "Constant"
	AttributeMappingSourceTypeFunction  AttributeMappingSourceType = "Function"
)

func PossibleValuesForAttributeMappingSourceType() []string {
	return []string{
		string(AttributeMappingSourceTypeAttribute),
		string(AttributeMappingSourceTypeConstant),
		string(AttributeMappingSourceTypeFunction),
	}
}

func parseAttributeMappingSourceType(input string) (*AttributeMappingSourceType, error) {
	vals := map[string]AttributeMappingSourceType{
		"attribute": AttributeMappingSourceTypeAttribute,
		"constant":  AttributeMappingSourceTypeConstant,
		"function":  AttributeMappingSourceTypeFunction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingSourceType(input)
	return &out, nil
}
