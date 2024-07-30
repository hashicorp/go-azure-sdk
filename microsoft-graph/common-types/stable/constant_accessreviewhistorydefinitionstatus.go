package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDefinitionStatus string

const (
	AccessReviewHistoryDefinitionStatus_Done       AccessReviewHistoryDefinitionStatus = "done"
	AccessReviewHistoryDefinitionStatus_Error      AccessReviewHistoryDefinitionStatus = "error"
	AccessReviewHistoryDefinitionStatus_Inprogress AccessReviewHistoryDefinitionStatus = "inprogress"
	AccessReviewHistoryDefinitionStatus_Requested  AccessReviewHistoryDefinitionStatus = "requested"
)

func PossibleValuesForAccessReviewHistoryDefinitionStatus() []string {
	return []string{
		string(AccessReviewHistoryDefinitionStatus_Done),
		string(AccessReviewHistoryDefinitionStatus_Error),
		string(AccessReviewHistoryDefinitionStatus_Inprogress),
		string(AccessReviewHistoryDefinitionStatus_Requested),
	}
}

func (s *AccessReviewHistoryDefinitionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewHistoryDefinitionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewHistoryDefinitionStatus(input string) (*AccessReviewHistoryDefinitionStatus, error) {
	vals := map[string]AccessReviewHistoryDefinitionStatus{
		"done":       AccessReviewHistoryDefinitionStatus_Done,
		"error":      AccessReviewHistoryDefinitionStatus_Error,
		"inprogress": AccessReviewHistoryDefinitionStatus_Inprogress,
		"requested":  AccessReviewHistoryDefinitionStatus_Requested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryDefinitionStatus(input)
	return &out, nil
}
