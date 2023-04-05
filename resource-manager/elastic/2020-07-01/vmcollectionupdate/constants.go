package vmcollectionupdate

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
