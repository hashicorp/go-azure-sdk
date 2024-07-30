package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestCategory string

const (
	FileAssessmentRequestCategory_Malware   FileAssessmentRequestCategory = "malware"
	FileAssessmentRequestCategory_Phishing  FileAssessmentRequestCategory = "phishing"
	FileAssessmentRequestCategory_Spam      FileAssessmentRequestCategory = "spam"
	FileAssessmentRequestCategory_Undefined FileAssessmentRequestCategory = "undefined"
)

func PossibleValuesForFileAssessmentRequestCategory() []string {
	return []string{
		string(FileAssessmentRequestCategory_Malware),
		string(FileAssessmentRequestCategory_Phishing),
		string(FileAssessmentRequestCategory_Spam),
		string(FileAssessmentRequestCategory_Undefined),
	}
}

func (s *FileAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileAssessmentRequestCategory(input string) (*FileAssessmentRequestCategory, error) {
	vals := map[string]FileAssessmentRequestCategory{
		"malware":   FileAssessmentRequestCategory_Malware,
		"phishing":  FileAssessmentRequestCategory_Phishing,
		"spam":      FileAssessmentRequestCategory_Spam,
		"undefined": FileAssessmentRequestCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestCategory(input)
	return &out, nil
}
