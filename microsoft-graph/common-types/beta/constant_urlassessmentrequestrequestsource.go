package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestRequestSource string

const (
	UrlAssessmentRequestRequestSource_Administrator UrlAssessmentRequestRequestSource = "administrator"
	UrlAssessmentRequestRequestSource_Undefined     UrlAssessmentRequestRequestSource = "undefined"
	UrlAssessmentRequestRequestSource_User          UrlAssessmentRequestRequestSource = "user"
)

func PossibleValuesForUrlAssessmentRequestRequestSource() []string {
	return []string{
		string(UrlAssessmentRequestRequestSource_Administrator),
		string(UrlAssessmentRequestRequestSource_Undefined),
		string(UrlAssessmentRequestRequestSource_User),
	}
}

func (s *UrlAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUrlAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUrlAssessmentRequestRequestSource(input string) (*UrlAssessmentRequestRequestSource, error) {
	vals := map[string]UrlAssessmentRequestRequestSource{
		"administrator": UrlAssessmentRequestRequestSource_Administrator,
		"undefined":     UrlAssessmentRequestRequestSource_Undefined,
		"user":          UrlAssessmentRequestRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestRequestSource(input)
	return &out, nil
}
