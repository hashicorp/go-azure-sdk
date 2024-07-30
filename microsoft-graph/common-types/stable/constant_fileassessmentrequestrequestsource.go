package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestRequestSource string

const (
	FileAssessmentRequestRequestSource_Administrator FileAssessmentRequestRequestSource = "administrator"
	FileAssessmentRequestRequestSource_Undefined     FileAssessmentRequestRequestSource = "undefined"
	FileAssessmentRequestRequestSource_User          FileAssessmentRequestRequestSource = "user"
)

func PossibleValuesForFileAssessmentRequestRequestSource() []string {
	return []string{
		string(FileAssessmentRequestRequestSource_Administrator),
		string(FileAssessmentRequestRequestSource_Undefined),
		string(FileAssessmentRequestRequestSource_User),
	}
}

func (s *FileAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileAssessmentRequestRequestSource(input string) (*FileAssessmentRequestRequestSource, error) {
	vals := map[string]FileAssessmentRequestRequestSource{
		"administrator": FileAssessmentRequestRequestSource_Administrator,
		"undefined":     FileAssessmentRequestRequestSource_Undefined,
		"user":          FileAssessmentRequestRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestRequestSource(input)
	return &out, nil
}
