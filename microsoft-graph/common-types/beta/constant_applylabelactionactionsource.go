package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyLabelActionActionSource string

const (
	ApplyLabelActionActionSource_Automatic   ApplyLabelActionActionSource = "automatic"
	ApplyLabelActionActionSource_Default     ApplyLabelActionActionSource = "default"
	ApplyLabelActionActionSource_Manual      ApplyLabelActionActionSource = "manual"
	ApplyLabelActionActionSource_Recommended ApplyLabelActionActionSource = "recommended"
)

func PossibleValuesForApplyLabelActionActionSource() []string {
	return []string{
		string(ApplyLabelActionActionSource_Automatic),
		string(ApplyLabelActionActionSource_Default),
		string(ApplyLabelActionActionSource_Manual),
		string(ApplyLabelActionActionSource_Recommended),
	}
}

func (s *ApplyLabelActionActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplyLabelActionActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplyLabelActionActionSource(input string) (*ApplyLabelActionActionSource, error) {
	vals := map[string]ApplyLabelActionActionSource{
		"automatic":   ApplyLabelActionActionSource_Automatic,
		"default":     ApplyLabelActionActionSource_Default,
		"manual":      ApplyLabelActionActionSource_Manual,
		"recommended": ApplyLabelActionActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplyLabelActionActionSource(input)
	return &out, nil
}
