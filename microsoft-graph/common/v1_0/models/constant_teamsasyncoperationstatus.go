package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperationStatus string

const (
	TeamsAsyncOperationStatusfailed     TeamsAsyncOperationStatus = "Failed"
	TeamsAsyncOperationStatusinProgress TeamsAsyncOperationStatus = "InProgress"
	TeamsAsyncOperationStatusinvalid    TeamsAsyncOperationStatus = "Invalid"
	TeamsAsyncOperationStatusnotStarted TeamsAsyncOperationStatus = "NotStarted"
	TeamsAsyncOperationStatussucceeded  TeamsAsyncOperationStatus = "Succeeded"
)

func PossibleValuesForTeamsAsyncOperationStatus() []string {
	return []string{
		string(TeamsAsyncOperationStatusfailed),
		string(TeamsAsyncOperationStatusinProgress),
		string(TeamsAsyncOperationStatusinvalid),
		string(TeamsAsyncOperationStatusnotStarted),
		string(TeamsAsyncOperationStatussucceeded),
	}
}

func parseTeamsAsyncOperationStatus(input string) (*TeamsAsyncOperationStatus, error) {
	vals := map[string]TeamsAsyncOperationStatus{
		"failed":     TeamsAsyncOperationStatusfailed,
		"inprogress": TeamsAsyncOperationStatusinProgress,
		"invalid":    TeamsAsyncOperationStatusinvalid,
		"notstarted": TeamsAsyncOperationStatusnotStarted,
		"succeeded":  TeamsAsyncOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationStatus(input)
	return &out, nil
}
