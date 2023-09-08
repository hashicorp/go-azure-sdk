package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataPolicyOperationStatus string

const (
	DataPolicyOperationStatuscomplete   DataPolicyOperationStatus = "Complete"
	DataPolicyOperationStatusfailed     DataPolicyOperationStatus = "Failed"
	DataPolicyOperationStatusnotStarted DataPolicyOperationStatus = "NotStarted"
	DataPolicyOperationStatusrunning    DataPolicyOperationStatus = "Running"
)

func PossibleValuesForDataPolicyOperationStatus() []string {
	return []string{
		string(DataPolicyOperationStatuscomplete),
		string(DataPolicyOperationStatusfailed),
		string(DataPolicyOperationStatusnotStarted),
		string(DataPolicyOperationStatusrunning),
	}
}

func parseDataPolicyOperationStatus(input string) (*DataPolicyOperationStatus, error) {
	vals := map[string]DataPolicyOperationStatus{
		"complete":   DataPolicyOperationStatuscomplete,
		"failed":     DataPolicyOperationStatusfailed,
		"notstarted": DataPolicyOperationStatusnotStarted,
		"running":    DataPolicyOperationStatusrunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataPolicyOperationStatus(input)
	return &out, nil
}
