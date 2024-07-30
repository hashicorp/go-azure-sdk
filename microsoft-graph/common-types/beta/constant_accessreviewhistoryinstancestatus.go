package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryInstanceStatus string

const (
	AccessReviewHistoryInstanceStatus_Done       AccessReviewHistoryInstanceStatus = "done"
	AccessReviewHistoryInstanceStatus_Error      AccessReviewHistoryInstanceStatus = "error"
	AccessReviewHistoryInstanceStatus_Inprogress AccessReviewHistoryInstanceStatus = "inprogress"
	AccessReviewHistoryInstanceStatus_Requested  AccessReviewHistoryInstanceStatus = "requested"
)

func PossibleValuesForAccessReviewHistoryInstanceStatus() []string {
	return []string{
		string(AccessReviewHistoryInstanceStatus_Done),
		string(AccessReviewHistoryInstanceStatus_Error),
		string(AccessReviewHistoryInstanceStatus_Inprogress),
		string(AccessReviewHistoryInstanceStatus_Requested),
	}
}

func (s *AccessReviewHistoryInstanceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewHistoryInstanceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewHistoryInstanceStatus(input string) (*AccessReviewHistoryInstanceStatus, error) {
	vals := map[string]AccessReviewHistoryInstanceStatus{
		"done":       AccessReviewHistoryInstanceStatus_Done,
		"error":      AccessReviewHistoryInstanceStatus_Error,
		"inprogress": AccessReviewHistoryInstanceStatus_Inprogress,
		"requested":  AccessReviewHistoryInstanceStatus_Requested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryInstanceStatus(input)
	return &out, nil
}
