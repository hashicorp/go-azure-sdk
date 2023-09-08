package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionOperationStatus string

const (
	ConnectionOperationStatuscompleted   ConnectionOperationStatus = "Completed"
	ConnectionOperationStatusfailed      ConnectionOperationStatus = "Failed"
	ConnectionOperationStatusinprogress  ConnectionOperationStatus = "Inprogress"
	ConnectionOperationStatusunspecified ConnectionOperationStatus = "Unspecified"
)

func PossibleValuesForConnectionOperationStatus() []string {
	return []string{
		string(ConnectionOperationStatuscompleted),
		string(ConnectionOperationStatusfailed),
		string(ConnectionOperationStatusinprogress),
		string(ConnectionOperationStatusunspecified),
	}
}

func parseConnectionOperationStatus(input string) (*ConnectionOperationStatus, error) {
	vals := map[string]ConnectionOperationStatus{
		"completed":   ConnectionOperationStatuscompleted,
		"failed":      ConnectionOperationStatusfailed,
		"inprogress":  ConnectionOperationStatusinprogress,
		"unspecified": ConnectionOperationStatusunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionOperationStatus(input)
	return &out, nil
}
