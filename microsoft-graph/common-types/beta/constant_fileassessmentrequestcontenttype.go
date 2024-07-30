package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestContentType string

const (
	FileAssessmentRequestContentType_File FileAssessmentRequestContentType = "file"
	FileAssessmentRequestContentType_Mail FileAssessmentRequestContentType = "mail"
	FileAssessmentRequestContentType_Url  FileAssessmentRequestContentType = "url"
)

func PossibleValuesForFileAssessmentRequestContentType() []string {
	return []string{
		string(FileAssessmentRequestContentType_File),
		string(FileAssessmentRequestContentType_Mail),
		string(FileAssessmentRequestContentType_Url),
	}
}

func (s *FileAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileAssessmentRequestContentType(input string) (*FileAssessmentRequestContentType, error) {
	vals := map[string]FileAssessmentRequestContentType{
		"file": FileAssessmentRequestContentType_File,
		"mail": FileAssessmentRequestContentType_Mail,
		"url":  FileAssessmentRequestContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestContentType(input)
	return &out, nil
}
