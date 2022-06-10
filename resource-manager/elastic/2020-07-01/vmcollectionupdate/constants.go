package vmcollectionupdate

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationName string

const (
	OperationNameAdd    OperationName = "Add"
	OperationNameDelete OperationName = "Delete"
)

func PossibleValuesForOperationName() []string {
	return []string{
		string(OperationNameAdd),
		string(OperationNameDelete),
	}
}

func parseOperationName(input string) (*OperationName, error) {
	vals := map[string]OperationName{
		"add":    OperationNameAdd,
		"delete": OperationNameDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationName(input)
	return &out, nil
}
