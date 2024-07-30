package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestContentType string

const (
	UrlAssessmentRequestContentType_File UrlAssessmentRequestContentType = "file"
	UrlAssessmentRequestContentType_Mail UrlAssessmentRequestContentType = "mail"
	UrlAssessmentRequestContentType_Url  UrlAssessmentRequestContentType = "url"
)

func PossibleValuesForUrlAssessmentRequestContentType() []string {
	return []string{
		string(UrlAssessmentRequestContentType_File),
		string(UrlAssessmentRequestContentType_Mail),
		string(UrlAssessmentRequestContentType_Url),
	}
}

func (s *UrlAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUrlAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUrlAssessmentRequestContentType(input string) (*UrlAssessmentRequestContentType, error) {
	vals := map[string]UrlAssessmentRequestContentType{
		"file": UrlAssessmentRequestContentType_File,
		"mail": UrlAssessmentRequestContentType_Mail,
		"url":  UrlAssessmentRequestContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestContentType(input)
	return &out, nil
}
