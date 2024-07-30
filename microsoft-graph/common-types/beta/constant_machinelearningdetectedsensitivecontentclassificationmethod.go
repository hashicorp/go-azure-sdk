package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentClassificationMethod string

const (
	MachineLearningDetectedSensitiveContentClassificationMethod_ExactDataMatch  MachineLearningDetectedSensitiveContentClassificationMethod = "exactDataMatch"
	MachineLearningDetectedSensitiveContentClassificationMethod_Fingerprint     MachineLearningDetectedSensitiveContentClassificationMethod = "fingerprint"
	MachineLearningDetectedSensitiveContentClassificationMethod_MachineLearning MachineLearningDetectedSensitiveContentClassificationMethod = "machineLearning"
	MachineLearningDetectedSensitiveContentClassificationMethod_PatternMatch    MachineLearningDetectedSensitiveContentClassificationMethod = "patternMatch"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentClassificationMethod() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentClassificationMethod_ExactDataMatch),
		string(MachineLearningDetectedSensitiveContentClassificationMethod_Fingerprint),
		string(MachineLearningDetectedSensitiveContentClassificationMethod_MachineLearning),
		string(MachineLearningDetectedSensitiveContentClassificationMethod_PatternMatch),
	}
}

func (s *MachineLearningDetectedSensitiveContentClassificationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineLearningDetectedSensitiveContentClassificationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineLearningDetectedSensitiveContentClassificationMethod(input string) (*MachineLearningDetectedSensitiveContentClassificationMethod, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentClassificationMethod{
		"exactdatamatch":  MachineLearningDetectedSensitiveContentClassificationMethod_ExactDataMatch,
		"fingerprint":     MachineLearningDetectedSensitiveContentClassificationMethod_Fingerprint,
		"machinelearning": MachineLearningDetectedSensitiveContentClassificationMethod_MachineLearning,
		"patternmatch":    MachineLearningDetectedSensitiveContentClassificationMethod_PatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentClassificationMethod(input)
	return &out, nil
}
