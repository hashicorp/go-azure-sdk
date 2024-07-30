package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestContentType string

const (
	MailAssessmentRequestContentType_File MailAssessmentRequestContentType = "file"
	MailAssessmentRequestContentType_Mail MailAssessmentRequestContentType = "mail"
	MailAssessmentRequestContentType_Url  MailAssessmentRequestContentType = "url"
)

func PossibleValuesForMailAssessmentRequestContentType() []string {
	return []string{
		string(MailAssessmentRequestContentType_File),
		string(MailAssessmentRequestContentType_Mail),
		string(MailAssessmentRequestContentType_Url),
	}
}

func (s *MailAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestContentType(input string) (*MailAssessmentRequestContentType, error) {
	vals := map[string]MailAssessmentRequestContentType{
		"file": MailAssessmentRequestContentType_File,
		"mail": MailAssessmentRequestContentType_Mail,
		"url":  MailAssessmentRequestContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestContentType(input)
	return &out, nil
}
