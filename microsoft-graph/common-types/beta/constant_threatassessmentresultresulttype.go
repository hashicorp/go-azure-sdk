package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentResultResultType string

const (
	ThreatAssessmentResultResultType_CheckPolicy ThreatAssessmentResultResultType = "checkPolicy"
	ThreatAssessmentResultResultType_Rescan      ThreatAssessmentResultResultType = "rescan"
)

func PossibleValuesForThreatAssessmentResultResultType() []string {
	return []string{
		string(ThreatAssessmentResultResultType_CheckPolicy),
		string(ThreatAssessmentResultResultType_Rescan),
	}
}

func (s *ThreatAssessmentResultResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentResultResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentResultResultType(input string) (*ThreatAssessmentResultResultType, error) {
	vals := map[string]ThreatAssessmentResultResultType{
		"checkpolicy": ThreatAssessmentResultResultType_CheckPolicy,
		"rescan":      ThreatAssessmentResultResultType_Rescan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentResultResultType(input)
	return &out, nil
}
