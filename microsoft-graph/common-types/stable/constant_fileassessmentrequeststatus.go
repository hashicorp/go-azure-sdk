package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestStatus string

const (
	FileAssessmentRequestStatus_Completed FileAssessmentRequestStatus = "completed"
	FileAssessmentRequestStatus_Pending   FileAssessmentRequestStatus = "pending"
)

func PossibleValuesForFileAssessmentRequestStatus() []string {
	return []string{
		string(FileAssessmentRequestStatus_Completed),
		string(FileAssessmentRequestStatus_Pending),
	}
}

func (s *FileAssessmentRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileAssessmentRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileAssessmentRequestStatus(input string) (*FileAssessmentRequestStatus, error) {
	vals := map[string]FileAssessmentRequestStatus{
		"completed": FileAssessmentRequestStatus_Completed,
		"pending":   FileAssessmentRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestStatus(input)
	return &out, nil
}
