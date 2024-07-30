package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentHeaderActionAlignment string

const (
	SecurityAddContentHeaderActionAlignment_Center SecurityAddContentHeaderActionAlignment = "center"
	SecurityAddContentHeaderActionAlignment_Left   SecurityAddContentHeaderActionAlignment = "left"
	SecurityAddContentHeaderActionAlignment_Right  SecurityAddContentHeaderActionAlignment = "right"
)

func PossibleValuesForSecurityAddContentHeaderActionAlignment() []string {
	return []string{
		string(SecurityAddContentHeaderActionAlignment_Center),
		string(SecurityAddContentHeaderActionAlignment_Left),
		string(SecurityAddContentHeaderActionAlignment_Right),
	}
}

func (s *SecurityAddContentHeaderActionAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAddContentHeaderActionAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAddContentHeaderActionAlignment(input string) (*SecurityAddContentHeaderActionAlignment, error) {
	vals := map[string]SecurityAddContentHeaderActionAlignment{
		"center": SecurityAddContentHeaderActionAlignment_Center,
		"left":   SecurityAddContentHeaderActionAlignment_Left,
		"right":  SecurityAddContentHeaderActionAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddContentHeaderActionAlignment(input)
	return &out, nil
}
