package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentScope string

const (
	MachineLearningDetectedSensitiveContentScope_FullDocument    MachineLearningDetectedSensitiveContentScope = "fullDocument"
	MachineLearningDetectedSensitiveContentScope_PartialDocument MachineLearningDetectedSensitiveContentScope = "partialDocument"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentScope() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentScope_FullDocument),
		string(MachineLearningDetectedSensitiveContentScope_PartialDocument),
	}
}

func (s *MachineLearningDetectedSensitiveContentScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMachineLearningDetectedSensitiveContentScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMachineLearningDetectedSensitiveContentScope(input string) (*MachineLearningDetectedSensitiveContentScope, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentScope{
		"fulldocument":    MachineLearningDetectedSensitiveContentScope_FullDocument,
		"partialdocument": MachineLearningDetectedSensitiveContentScope_PartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentScope(input)
	return &out, nil
}
