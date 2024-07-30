package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddHeaderAlignment string

const (
	AddHeaderAlignment_Center AddHeaderAlignment = "center"
	AddHeaderAlignment_Left   AddHeaderAlignment = "left"
	AddHeaderAlignment_Right  AddHeaderAlignment = "right"
)

func PossibleValuesForAddHeaderAlignment() []string {
	return []string{
		string(AddHeaderAlignment_Center),
		string(AddHeaderAlignment_Left),
		string(AddHeaderAlignment_Right),
	}
}

func (s *AddHeaderAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddHeaderAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddHeaderAlignment(input string) (*AddHeaderAlignment, error) {
	vals := map[string]AddHeaderAlignment{
		"center": AddHeaderAlignment_Center,
		"left":   AddHeaderAlignment_Left,
		"right":  AddHeaderAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddHeaderAlignment(input)
	return &out, nil
}
