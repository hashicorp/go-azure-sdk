package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConnectionOperationStatus string

const (
	ExternalConnectorsConnectionOperationStatuscompleted   ExternalConnectorsConnectionOperationStatus = "Completed"
	ExternalConnectorsConnectionOperationStatusfailed      ExternalConnectorsConnectionOperationStatus = "Failed"
	ExternalConnectorsConnectionOperationStatusinprogress  ExternalConnectorsConnectionOperationStatus = "Inprogress"
	ExternalConnectorsConnectionOperationStatusunspecified ExternalConnectorsConnectionOperationStatus = "Unspecified"
)

func PossibleValuesForExternalConnectorsConnectionOperationStatus() []string {
	return []string{
		string(ExternalConnectorsConnectionOperationStatuscompleted),
		string(ExternalConnectorsConnectionOperationStatusfailed),
		string(ExternalConnectorsConnectionOperationStatusinprogress),
		string(ExternalConnectorsConnectionOperationStatusunspecified),
	}
}

func parseExternalConnectorsConnectionOperationStatus(input string) (*ExternalConnectorsConnectionOperationStatus, error) {
	vals := map[string]ExternalConnectorsConnectionOperationStatus{
		"completed":   ExternalConnectorsConnectionOperationStatuscompleted,
		"failed":      ExternalConnectorsConnectionOperationStatusfailed,
		"inprogress":  ExternalConnectorsConnectionOperationStatusinprogress,
		"unspecified": ExternalConnectorsConnectionOperationStatusunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsConnectionOperationStatus(input)
	return &out, nil
}
