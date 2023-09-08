package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingFlowType string

const (
	AttributeMappingFlowTypeAlways            AttributeMappingFlowType = "Always"
	AttributeMappingFlowTypeAttributeAddOnly  AttributeMappingFlowType = "AttributeAddOnly"
	AttributeMappingFlowTypeMultiValueAddOnly AttributeMappingFlowType = "MultiValueAddOnly"
	AttributeMappingFlowTypeObjectAddOnly     AttributeMappingFlowType = "ObjectAddOnly"
	AttributeMappingFlowTypeValueAddOnly      AttributeMappingFlowType = "ValueAddOnly"
)

func PossibleValuesForAttributeMappingFlowType() []string {
	return []string{
		string(AttributeMappingFlowTypeAlways),
		string(AttributeMappingFlowTypeAttributeAddOnly),
		string(AttributeMappingFlowTypeMultiValueAddOnly),
		string(AttributeMappingFlowTypeObjectAddOnly),
		string(AttributeMappingFlowTypeValueAddOnly),
	}
}

func parseAttributeMappingFlowType(input string) (*AttributeMappingFlowType, error) {
	vals := map[string]AttributeMappingFlowType{
		"always":            AttributeMappingFlowTypeAlways,
		"attributeaddonly":  AttributeMappingFlowTypeAttributeAddOnly,
		"multivalueaddonly": AttributeMappingFlowTypeMultiValueAddOnly,
		"objectaddonly":     AttributeMappingFlowTypeObjectAddOnly,
		"valueaddonly":      AttributeMappingFlowTypeValueAddOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingFlowType(input)
	return &out, nil
}
