package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentMatchTolerance string

const (
	MachineLearningDetectedSensitiveContentMatchTolerance_Exact MachineLearningDetectedSensitiveContentMatchTolerance = "exact"
	MachineLearningDetectedSensitiveContentMatchTolerance_Near  MachineLearningDetectedSensitiveContentMatchTolerance = "near"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentMatchTolerance() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentMatchTolerance_Exact),
		string(MachineLearningDetectedSensitiveContentMatchTolerance_Near),
	}
}

func (s *MachineLearningDetectedSensitiveContentMatchTolerance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineLearningDetectedSensitiveContentMatchTolerance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineLearningDetectedSensitiveContentMatchTolerance(input string) (*MachineLearningDetectedSensitiveContentMatchTolerance, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentMatchTolerance{
		"exact": MachineLearningDetectedSensitiveContentMatchTolerance_Exact,
		"near":  MachineLearningDetectedSensitiveContentMatchTolerance_Near,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentMatchTolerance(input)
	return &out, nil
}
