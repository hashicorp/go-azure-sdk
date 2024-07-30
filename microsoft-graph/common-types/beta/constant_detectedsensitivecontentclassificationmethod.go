package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentClassificationMethod string

const (
	DetectedSensitiveContentClassificationMethod_ExactDataMatch  DetectedSensitiveContentClassificationMethod = "exactDataMatch"
	DetectedSensitiveContentClassificationMethod_Fingerprint     DetectedSensitiveContentClassificationMethod = "fingerprint"
	DetectedSensitiveContentClassificationMethod_MachineLearning DetectedSensitiveContentClassificationMethod = "machineLearning"
	DetectedSensitiveContentClassificationMethod_PatternMatch    DetectedSensitiveContentClassificationMethod = "patternMatch"
)

func PossibleValuesForDetectedSensitiveContentClassificationMethod() []string {
	return []string{
		string(DetectedSensitiveContentClassificationMethod_ExactDataMatch),
		string(DetectedSensitiveContentClassificationMethod_Fingerprint),
		string(DetectedSensitiveContentClassificationMethod_MachineLearning),
		string(DetectedSensitiveContentClassificationMethod_PatternMatch),
	}
}

func (s *DetectedSensitiveContentClassificationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedSensitiveContentClassificationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedSensitiveContentClassificationMethod(input string) (*DetectedSensitiveContentClassificationMethod, error) {
	vals := map[string]DetectedSensitiveContentClassificationMethod{
		"exactdatamatch":  DetectedSensitiveContentClassificationMethod_ExactDataMatch,
		"fingerprint":     DetectedSensitiveContentClassificationMethod_Fingerprint,
		"machinelearning": DetectedSensitiveContentClassificationMethod_MachineLearning,
		"patternmatch":    DetectedSensitiveContentClassificationMethod_PatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentClassificationMethod(input)
	return &out, nil
}
