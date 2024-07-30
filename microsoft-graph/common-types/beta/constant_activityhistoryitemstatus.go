package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItemStatus string

const (
	ActivityHistoryItemStatus_Active  ActivityHistoryItemStatus = "active"
	ActivityHistoryItemStatus_Deleted ActivityHistoryItemStatus = "deleted"
	ActivityHistoryItemStatus_Ignored ActivityHistoryItemStatus = "ignored"
	ActivityHistoryItemStatus_Updated ActivityHistoryItemStatus = "updated"
)

func PossibleValuesForActivityHistoryItemStatus() []string {
	return []string{
		string(ActivityHistoryItemStatus_Active),
		string(ActivityHistoryItemStatus_Deleted),
		string(ActivityHistoryItemStatus_Ignored),
		string(ActivityHistoryItemStatus_Updated),
	}
}

func (s *ActivityHistoryItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityHistoryItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityHistoryItemStatus(input string) (*ActivityHistoryItemStatus, error) {
	vals := map[string]ActivityHistoryItemStatus{
		"active":  ActivityHistoryItemStatus_Active,
		"deleted": ActivityHistoryItemStatus_Deleted,
		"ignored": ActivityHistoryItemStatus_Ignored,
		"updated": ActivityHistoryItemStatus_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityHistoryItemStatus(input)
	return &out, nil
}
