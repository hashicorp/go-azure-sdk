package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityApplyLabelActionActionSource string

const (
	SecurityApplyLabelActionActionSource_Automatic   SecurityApplyLabelActionActionSource = "automatic"
	SecurityApplyLabelActionActionSource_Default     SecurityApplyLabelActionActionSource = "default"
	SecurityApplyLabelActionActionSource_Manual      SecurityApplyLabelActionActionSource = "manual"
	SecurityApplyLabelActionActionSource_Recommended SecurityApplyLabelActionActionSource = "recommended"
)

func PossibleValuesForSecurityApplyLabelActionActionSource() []string {
	return []string{
		string(SecurityApplyLabelActionActionSource_Automatic),
		string(SecurityApplyLabelActionActionSource_Default),
		string(SecurityApplyLabelActionActionSource_Manual),
		string(SecurityApplyLabelActionActionSource_Recommended),
	}
}

func (s *SecurityApplyLabelActionActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityApplyLabelActionActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityApplyLabelActionActionSource(input string) (*SecurityApplyLabelActionActionSource, error) {
	vals := map[string]SecurityApplyLabelActionActionSource{
		"automatic":   SecurityApplyLabelActionActionSource_Automatic,
		"default":     SecurityApplyLabelActionActionSource_Default,
		"manual":      SecurityApplyLabelActionActionSource_Manual,
		"recommended": SecurityApplyLabelActionActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityApplyLabelActionActionSource(input)
	return &out, nil
}
