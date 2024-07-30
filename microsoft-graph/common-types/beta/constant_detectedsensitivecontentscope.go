package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentScope string

const (
	DetectedSensitiveContentScope_FullDocument    DetectedSensitiveContentScope = "fullDocument"
	DetectedSensitiveContentScope_PartialDocument DetectedSensitiveContentScope = "partialDocument"
)

func PossibleValuesForDetectedSensitiveContentScope() []string {
	return []string{
		string(DetectedSensitiveContentScope_FullDocument),
		string(DetectedSensitiveContentScope_PartialDocument),
	}
}

func (s *DetectedSensitiveContentScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedSensitiveContentScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedSensitiveContentScope(input string) (*DetectedSensitiveContentScope, error) {
	vals := map[string]DetectedSensitiveContentScope{
		"fulldocument":    DetectedSensitiveContentScope_FullDocument,
		"partialdocument": DetectedSensitiveContentScope_PartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentScope(input)
	return &out, nil
}
