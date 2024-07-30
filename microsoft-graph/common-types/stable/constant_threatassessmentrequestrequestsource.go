package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestRequestSource string

const (
	ThreatAssessmentRequestRequestSource_Administrator ThreatAssessmentRequestRequestSource = "administrator"
	ThreatAssessmentRequestRequestSource_Undefined     ThreatAssessmentRequestRequestSource = "undefined"
	ThreatAssessmentRequestRequestSource_User          ThreatAssessmentRequestRequestSource = "user"
)

func PossibleValuesForThreatAssessmentRequestRequestSource() []string {
	return []string{
		string(ThreatAssessmentRequestRequestSource_Administrator),
		string(ThreatAssessmentRequestRequestSource_Undefined),
		string(ThreatAssessmentRequestRequestSource_User),
	}
}

func (s *ThreatAssessmentRequestRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestRequestSource(input string) (*ThreatAssessmentRequestRequestSource, error) {
	vals := map[string]ThreatAssessmentRequestRequestSource{
		"administrator": ThreatAssessmentRequestRequestSource_Administrator,
		"undefined":     ThreatAssessmentRequestRequestSource_Undefined,
		"user":          ThreatAssessmentRequestRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestRequestSource(input)
	return &out, nil
}
