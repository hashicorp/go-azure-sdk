package privateendpoint

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusValues string

const (
	OperationStatusValuesCanceled   OperationStatusValues = "Canceled"
	OperationStatusValuesFailed     OperationStatusValues = "Failed"
	OperationStatusValuesInProgress OperationStatusValues = "InProgress"
	OperationStatusValuesInvalid    OperationStatusValues = "Invalid"
	OperationStatusValuesSucceeded  OperationStatusValues = "Succeeded"
)

func PossibleValuesForOperationStatusValues() []string {
	return []string{
		string(OperationStatusValuesCanceled),
		string(OperationStatusValuesFailed),
		string(OperationStatusValuesInProgress),
		string(OperationStatusValuesInvalid),
		string(OperationStatusValuesSucceeded),
	}
}

func parseOperationStatusValues(input string) (*OperationStatusValues, error) {
	vals := map[string]OperationStatusValues{
		"canceled":   OperationStatusValuesCanceled,
		"failed":     OperationStatusValuesFailed,
		"inprogress": OperationStatusValuesInProgress,
		"invalid":    OperationStatusValuesInvalid,
		"succeeded":  OperationStatusValuesSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusValues(input)
	return &out, nil
}
