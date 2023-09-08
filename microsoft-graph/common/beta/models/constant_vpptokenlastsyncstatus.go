package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenLastSyncStatus string

const (
	VppTokenLastSyncStatuscompleted  VppTokenLastSyncStatus = "Completed"
	VppTokenLastSyncStatusfailed     VppTokenLastSyncStatus = "Failed"
	VppTokenLastSyncStatusinProgress VppTokenLastSyncStatus = "InProgress"
	VppTokenLastSyncStatusnone       VppTokenLastSyncStatus = "None"
)

func PossibleValuesForVppTokenLastSyncStatus() []string {
	return []string{
		string(VppTokenLastSyncStatuscompleted),
		string(VppTokenLastSyncStatusfailed),
		string(VppTokenLastSyncStatusinProgress),
		string(VppTokenLastSyncStatusnone),
	}
}

func parseVppTokenLastSyncStatus(input string) (*VppTokenLastSyncStatus, error) {
	vals := map[string]VppTokenLastSyncStatus{
		"completed":  VppTokenLastSyncStatuscompleted,
		"failed":     VppTokenLastSyncStatusfailed,
		"inprogress": VppTokenLastSyncStatusinProgress,
		"none":       VppTokenLastSyncStatusnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenLastSyncStatus(input)
	return &out, nil
}
