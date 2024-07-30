package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationOverrideClassifyAs string

const (
	InferenceClassificationOverrideClassifyAs_Focused InferenceClassificationOverrideClassifyAs = "focused"
	InferenceClassificationOverrideClassifyAs_Other   InferenceClassificationOverrideClassifyAs = "other"
)

func PossibleValuesForInferenceClassificationOverrideClassifyAs() []string {
	return []string{
		string(InferenceClassificationOverrideClassifyAs_Focused),
		string(InferenceClassificationOverrideClassifyAs_Other),
	}
}

func (s *InferenceClassificationOverrideClassifyAs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInferenceClassificationOverrideClassifyAs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInferenceClassificationOverrideClassifyAs(input string) (*InferenceClassificationOverrideClassifyAs, error) {
	vals := map[string]InferenceClassificationOverrideClassifyAs{
		"focused": InferenceClassificationOverrideClassifyAs_Focused,
		"other":   InferenceClassificationOverrideClassifyAs_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InferenceClassificationOverrideClassifyAs(input)
	return &out, nil
}
