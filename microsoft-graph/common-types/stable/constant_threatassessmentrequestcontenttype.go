package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestContentType string

const (
	ThreatAssessmentRequestContentType_File ThreatAssessmentRequestContentType = "file"
	ThreatAssessmentRequestContentType_Mail ThreatAssessmentRequestContentType = "mail"
	ThreatAssessmentRequestContentType_Url  ThreatAssessmentRequestContentType = "url"
)

func PossibleValuesForThreatAssessmentRequestContentType() []string {
	return []string{
		string(ThreatAssessmentRequestContentType_File),
		string(ThreatAssessmentRequestContentType_Mail),
		string(ThreatAssessmentRequestContentType_Url),
	}
}

func (s *ThreatAssessmentRequestContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestContentType(input string) (*ThreatAssessmentRequestContentType, error) {
	vals := map[string]ThreatAssessmentRequestContentType{
		"file": ThreatAssessmentRequestContentType_File,
		"mail": ThreatAssessmentRequestContentType_Mail,
		"url":  ThreatAssessmentRequestContentType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestContentType(input)
	return &out, nil
}
