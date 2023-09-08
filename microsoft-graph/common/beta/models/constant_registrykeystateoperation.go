package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryKeyStateOperation string

const (
	RegistryKeyStateOperationcreate  RegistryKeyStateOperation = "Create"
	RegistryKeyStateOperationdelete  RegistryKeyStateOperation = "Delete"
	RegistryKeyStateOperationmodify  RegistryKeyStateOperation = "Modify"
	RegistryKeyStateOperationunknown RegistryKeyStateOperation = "Unknown"
)

func PossibleValuesForRegistryKeyStateOperation() []string {
	return []string{
		string(RegistryKeyStateOperationcreate),
		string(RegistryKeyStateOperationdelete),
		string(RegistryKeyStateOperationmodify),
		string(RegistryKeyStateOperationunknown),
	}
}

func parseRegistryKeyStateOperation(input string) (*RegistryKeyStateOperation, error) {
	vals := map[string]RegistryKeyStateOperation{
		"create":  RegistryKeyStateOperationcreate,
		"delete":  RegistryKeyStateOperationdelete,
		"modify":  RegistryKeyStateOperationmodify,
		"unknown": RegistryKeyStateOperationunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryKeyStateOperation(input)
	return &out, nil
}
