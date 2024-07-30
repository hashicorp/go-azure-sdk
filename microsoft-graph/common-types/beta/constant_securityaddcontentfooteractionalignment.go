package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentFooterActionAlignment string

const (
	SecurityAddContentFooterActionAlignment_Center SecurityAddContentFooterActionAlignment = "center"
	SecurityAddContentFooterActionAlignment_Left   SecurityAddContentFooterActionAlignment = "left"
	SecurityAddContentFooterActionAlignment_Right  SecurityAddContentFooterActionAlignment = "right"
)

func PossibleValuesForSecurityAddContentFooterActionAlignment() []string {
	return []string{
		string(SecurityAddContentFooterActionAlignment_Center),
		string(SecurityAddContentFooterActionAlignment_Left),
		string(SecurityAddContentFooterActionAlignment_Right),
	}
}

func (s *SecurityAddContentFooterActionAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAddContentFooterActionAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAddContentFooterActionAlignment(input string) (*SecurityAddContentFooterActionAlignment, error) {
	vals := map[string]SecurityAddContentFooterActionAlignment{
		"center": SecurityAddContentFooterActionAlignment_Center,
		"left":   SecurityAddContentFooterActionAlignment_Left,
		"right":  SecurityAddContentFooterActionAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddContentFooterActionAlignment(input)
	return &out, nil
}
