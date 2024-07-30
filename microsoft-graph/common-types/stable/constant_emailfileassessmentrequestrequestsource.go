package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestRequestSource string

const (
	EmailFileAssessmentRequestRequestSource_Administrator EmailFileAssessmentRequestRequestSource = "administrator"
	EmailFileAssessmentRequestRequestSource_Undefined     EmailFileAssessmentRequestRequestSource = "undefined"
	EmailFileAssessmentRequestRequestSource_User          EmailFileAssessmentRequestRequestSource = "user"
)

func PossibleValuesForEmailFileAssessmentRequestRequestSource() []string {
	return []string{
		string(EmailFileAssessmentRequestRequestSource_Administrator),
		string(EmailFileAssessmentRequestRequestSource_Undefined),
		string(EmailFileAssessmentRequestRequestSource_User),
	}
}

func (s *EmailFileAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailFileAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailFileAssessmentRequestRequestSource(input string) (*EmailFileAssessmentRequestRequestSource, error) {
	vals := map[string]EmailFileAssessmentRequestRequestSource{
		"administrator": EmailFileAssessmentRequestRequestSource_Administrator,
		"undefined":     EmailFileAssessmentRequestRequestSource_Undefined,
		"user":          EmailFileAssessmentRequestRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestRequestSource(input)
	return &out, nil
}
