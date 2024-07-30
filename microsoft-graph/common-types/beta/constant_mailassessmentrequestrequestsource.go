package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestRequestSource string

const (
	MailAssessmentRequestRequestSource_Administrator MailAssessmentRequestRequestSource = "administrator"
	MailAssessmentRequestRequestSource_Undefined     MailAssessmentRequestRequestSource = "undefined"
	MailAssessmentRequestRequestSource_User          MailAssessmentRequestRequestSource = "user"
)

func PossibleValuesForMailAssessmentRequestRequestSource() []string {
	return []string{
		string(MailAssessmentRequestRequestSource_Administrator),
		string(MailAssessmentRequestRequestSource_Undefined),
		string(MailAssessmentRequestRequestSource_User),
	}
}

func (s *MailAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestRequestSource(input string) (*MailAssessmentRequestRequestSource, error) {
	vals := map[string]MailAssessmentRequestRequestSource{
		"administrator": MailAssessmentRequestRequestSource_Administrator,
		"undefined":     MailAssessmentRequestRequestSource_Undefined,
		"user":          MailAssessmentRequestRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestRequestSource(input)
	return &out, nil
}
