package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestContentType string

const (
	EmailFileAssessmentRequestContentType_File EmailFileAssessmentRequestContentType = "file"
	EmailFileAssessmentRequestContentType_Mail EmailFileAssessmentRequestContentType = "mail"
	EmailFileAssessmentRequestContentType_Url  EmailFileAssessmentRequestContentType = "url"
)

func PossibleValuesForEmailFileAssessmentRequestContentType() []string {
	return []string{
		string(EmailFileAssessmentRequestContentType_File),
		string(EmailFileAssessmentRequestContentType_Mail),
		string(EmailFileAssessmentRequestContentType_Url),
	}
}

func (s *EmailFileAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailFileAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailFileAssessmentRequestContentType(input string) (*EmailFileAssessmentRequestContentType, error) {
	vals := map[string]EmailFileAssessmentRequestContentType{
		"file": EmailFileAssessmentRequestContentType_File,
		"mail": EmailFileAssessmentRequestContentType_Mail,
		"url":  EmailFileAssessmentRequestContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestContentType(input)
	return &out, nil
}
