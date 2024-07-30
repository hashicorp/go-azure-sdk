package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddFooterAlignment string

const (
	AddFooterAlignment_Center AddFooterAlignment = "center"
	AddFooterAlignment_Left   AddFooterAlignment = "left"
	AddFooterAlignment_Right  AddFooterAlignment = "right"
)

func PossibleValuesForAddFooterAlignment() []string {
	return []string{
		string(AddFooterAlignment_Center),
		string(AddFooterAlignment_Left),
		string(AddFooterAlignment_Right),
	}
}

func (s *AddFooterAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddFooterAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddFooterAlignment(input string) (*AddFooterAlignment, error) {
	vals := map[string]AddFooterAlignment{
		"center": AddFooterAlignment_Center,
		"left":   AddFooterAlignment_Left,
		"right":  AddFooterAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddFooterAlignment(input)
	return &out, nil
}
