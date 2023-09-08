package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobExportJobStatus string

const (
	CloudPCExportJobExportJobStatuscompleted  CloudPCExportJobExportJobStatus = "Completed"
	CloudPCExportJobExportJobStatusfailed     CloudPCExportJobExportJobStatus = "Failed"
	CloudPCExportJobExportJobStatusinProgress CloudPCExportJobExportJobStatus = "InProgress"
	CloudPCExportJobExportJobStatusnotStarted CloudPCExportJobExportJobStatus = "NotStarted"
)

func PossibleValuesForCloudPCExportJobExportJobStatus() []string {
	return []string{
		string(CloudPCExportJobExportJobStatuscompleted),
		string(CloudPCExportJobExportJobStatusfailed),
		string(CloudPCExportJobExportJobStatusinProgress),
		string(CloudPCExportJobExportJobStatusnotStarted),
	}
}

func parseCloudPCExportJobExportJobStatus(input string) (*CloudPCExportJobExportJobStatus, error) {
	vals := map[string]CloudPCExportJobExportJobStatus{
		"completed":  CloudPCExportJobExportJobStatuscompleted,
		"failed":     CloudPCExportJobExportJobStatusfailed,
		"inprogress": CloudPCExportJobExportJobStatusinProgress,
		"notstarted": CloudPCExportJobExportJobStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExportJobExportJobStatus(input)
	return &out, nil
}
