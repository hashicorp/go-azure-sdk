package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentInfoState string

const (
	ContentInfoState_Motion ContentInfoState = "motion"
	ContentInfoState_Rest   ContentInfoState = "rest"
	ContentInfoState_Use    ContentInfoState = "use"
)

func PossibleValuesForContentInfoState() []string {
	return []string{
		string(ContentInfoState_Motion),
		string(ContentInfoState_Rest),
		string(ContentInfoState_Use),
	}
}

func (s *ContentInfoState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentInfoState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentInfoState(input string) (*ContentInfoState, error) {
	vals := map[string]ContentInfoState{
		"motion": ContentInfoState_Motion,
		"rest":   ContentInfoState_Rest,
		"use":    ContentInfoState_Use,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentInfoState(input)
	return &out, nil
}
