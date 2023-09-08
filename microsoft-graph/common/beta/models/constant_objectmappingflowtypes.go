package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingFlowTypes string

const (
	ObjectMappingFlowTypesAdd    ObjectMappingFlowTypes = "Add"
	ObjectMappingFlowTypesDelete ObjectMappingFlowTypes = "Delete"
	ObjectMappingFlowTypesNone   ObjectMappingFlowTypes = "None"
	ObjectMappingFlowTypesUpdate ObjectMappingFlowTypes = "Update"
)

func PossibleValuesForObjectMappingFlowTypes() []string {
	return []string{
		string(ObjectMappingFlowTypesAdd),
		string(ObjectMappingFlowTypesDelete),
		string(ObjectMappingFlowTypesNone),
		string(ObjectMappingFlowTypesUpdate),
	}
}

func parseObjectMappingFlowTypes(input string) (*ObjectMappingFlowTypes, error) {
	vals := map[string]ObjectMappingFlowTypes{
		"add":    ObjectMappingFlowTypesAdd,
		"delete": ObjectMappingFlowTypesDelete,
		"none":   ObjectMappingFlowTypesNone,
		"update": ObjectMappingFlowTypesUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObjectMappingFlowTypes(input)
	return &out, nil
}
