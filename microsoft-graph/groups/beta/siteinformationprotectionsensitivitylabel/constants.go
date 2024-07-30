package siteinformationprotectionsensitivitylabel

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentLabelApplicationMode string

const (
	CurrentLabelApplicationModeAutomatic   CurrentLabelApplicationMode = "automatic"
	CurrentLabelApplicationModeManual      CurrentLabelApplicationMode = "manual"
	CurrentLabelApplicationModeRecommended CurrentLabelApplicationMode = "recommended"
)

func PossibleValuesForCurrentLabelApplicationMode() []string {
	return []string{
		string(CurrentLabelApplicationModeAutomatic),
		string(CurrentLabelApplicationModeManual),
		string(CurrentLabelApplicationModeRecommended),
	}
}

func (s *CurrentLabelApplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCurrentLabelApplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCurrentLabelApplicationMode(input string) (*CurrentLabelApplicationMode, error) {
	vals := map[string]CurrentLabelApplicationMode{
		"automatic":   CurrentLabelApplicationModeAutomatic,
		"manual":      CurrentLabelApplicationModeManual,
		"recommended": CurrentLabelApplicationModeRecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CurrentLabelApplicationMode(input)
	return &out, nil
}
