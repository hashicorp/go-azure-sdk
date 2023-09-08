package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItemStatus string

const (
	ActivityHistoryItemStatusactive  ActivityHistoryItemStatus = "Active"
	ActivityHistoryItemStatusdeleted ActivityHistoryItemStatus = "Deleted"
	ActivityHistoryItemStatusignored ActivityHistoryItemStatus = "Ignored"
	ActivityHistoryItemStatusupdated ActivityHistoryItemStatus = "Updated"
)

func PossibleValuesForActivityHistoryItemStatus() []string {
	return []string{
		string(ActivityHistoryItemStatusactive),
		string(ActivityHistoryItemStatusdeleted),
		string(ActivityHistoryItemStatusignored),
		string(ActivityHistoryItemStatusupdated),
	}
}

func parseActivityHistoryItemStatus(input string) (*ActivityHistoryItemStatus, error) {
	vals := map[string]ActivityHistoryItemStatus{
		"active":  ActivityHistoryItemStatusactive,
		"deleted": ActivityHistoryItemStatusdeleted,
		"ignored": ActivityHistoryItemStatusignored,
		"updated": ActivityHistoryItemStatusupdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityHistoryItemStatus(input)
	return &out, nil
}
