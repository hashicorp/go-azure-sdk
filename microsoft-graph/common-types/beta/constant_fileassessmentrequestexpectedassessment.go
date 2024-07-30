package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestExpectedAssessment string

const (
	FileAssessmentRequestExpectedAssessment_Block   FileAssessmentRequestExpectedAssessment = "block"
	FileAssessmentRequestExpectedAssessment_Unblock FileAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForFileAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(FileAssessmentRequestExpectedAssessment_Block),
		string(FileAssessmentRequestExpectedAssessment_Unblock),
	}
}

func (s *FileAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileAssessmentRequestExpectedAssessment(input string) (*FileAssessmentRequestExpectedAssessment, error) {
	vals := map[string]FileAssessmentRequestExpectedAssessment{
		"block":   FileAssessmentRequestExpectedAssessment_Block,
		"unblock": FileAssessmentRequestExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
