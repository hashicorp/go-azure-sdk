package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeClassificationMethod string

const (
	SensitiveTypeClassificationMethod_ExactDataMatch  SensitiveTypeClassificationMethod = "exactDataMatch"
	SensitiveTypeClassificationMethod_Fingerprint     SensitiveTypeClassificationMethod = "fingerprint"
	SensitiveTypeClassificationMethod_MachineLearning SensitiveTypeClassificationMethod = "machineLearning"
	SensitiveTypeClassificationMethod_PatternMatch    SensitiveTypeClassificationMethod = "patternMatch"
)

func PossibleValuesForSensitiveTypeClassificationMethod() []string {
	return []string{
		string(SensitiveTypeClassificationMethod_ExactDataMatch),
		string(SensitiveTypeClassificationMethod_Fingerprint),
		string(SensitiveTypeClassificationMethod_MachineLearning),
		string(SensitiveTypeClassificationMethod_PatternMatch),
	}
}

func (s *SensitiveTypeClassificationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitiveTypeClassificationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitiveTypeClassificationMethod(input string) (*SensitiveTypeClassificationMethod, error) {
	vals := map[string]SensitiveTypeClassificationMethod{
		"exactdatamatch":  SensitiveTypeClassificationMethod_ExactDataMatch,
		"fingerprint":     SensitiveTypeClassificationMethod_Fingerprint,
		"machinelearning": SensitiveTypeClassificationMethod_MachineLearning,
		"patternmatch":    SensitiveTypeClassificationMethod_PatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeClassificationMethod(input)
	return &out, nil
}
