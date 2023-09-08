package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkRestoreTimeRange string

const (
	CloudPCBulkRestoreTimeRangeafter         CloudPCBulkRestoreTimeRange = "After"
	CloudPCBulkRestoreTimeRangebefore        CloudPCBulkRestoreTimeRange = "Before"
	CloudPCBulkRestoreTimeRangebeforeOrAfter CloudPCBulkRestoreTimeRange = "BeforeOrAfter"
)

func PossibleValuesForCloudPCBulkRestoreTimeRange() []string {
	return []string{
		string(CloudPCBulkRestoreTimeRangeafter),
		string(CloudPCBulkRestoreTimeRangebefore),
		string(CloudPCBulkRestoreTimeRangebeforeOrAfter),
	}
}

func parseCloudPCBulkRestoreTimeRange(input string) (*CloudPCBulkRestoreTimeRange, error) {
	vals := map[string]CloudPCBulkRestoreTimeRange{
		"after":         CloudPCBulkRestoreTimeRangeafter,
		"before":        CloudPCBulkRestoreTimeRangebefore,
		"beforeorafter": CloudPCBulkRestoreTimeRangebeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCBulkRestoreTimeRange(input)
	return &out, nil
}
