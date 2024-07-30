package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkRestoreTimeRange string

const (
	CloudPCBulkRestoreTimeRange_After         CloudPCBulkRestoreTimeRange = "after"
	CloudPCBulkRestoreTimeRange_Before        CloudPCBulkRestoreTimeRange = "before"
	CloudPCBulkRestoreTimeRange_BeforeOrAfter CloudPCBulkRestoreTimeRange = "beforeOrAfter"
)

func PossibleValuesForCloudPCBulkRestoreTimeRange() []string {
	return []string{
		string(CloudPCBulkRestoreTimeRange_After),
		string(CloudPCBulkRestoreTimeRange_Before),
		string(CloudPCBulkRestoreTimeRange_BeforeOrAfter),
	}
}

func (s *CloudPCBulkRestoreTimeRange) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCBulkRestoreTimeRange(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCBulkRestoreTimeRange(input string) (*CloudPCBulkRestoreTimeRange, error) {
	vals := map[string]CloudPCBulkRestoreTimeRange{
		"after":         CloudPCBulkRestoreTimeRange_After,
		"before":        CloudPCBulkRestoreTimeRange_Before,
		"beforeorafter": CloudPCBulkRestoreTimeRange_BeforeOrAfter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCBulkRestoreTimeRange(input)
	return &out, nil
}
