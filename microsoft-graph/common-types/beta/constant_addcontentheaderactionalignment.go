package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddContentHeaderActionAlignment string

const (
	AddContentHeaderActionAlignment_Center AddContentHeaderActionAlignment = "center"
	AddContentHeaderActionAlignment_Left   AddContentHeaderActionAlignment = "left"
	AddContentHeaderActionAlignment_Right  AddContentHeaderActionAlignment = "right"
)

func PossibleValuesForAddContentHeaderActionAlignment() []string {
	return []string{
		string(AddContentHeaderActionAlignment_Center),
		string(AddContentHeaderActionAlignment_Left),
		string(AddContentHeaderActionAlignment_Right),
	}
}

func (s *AddContentHeaderActionAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddContentHeaderActionAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddContentHeaderActionAlignment(input string) (*AddContentHeaderActionAlignment, error) {
	vals := map[string]AddContentHeaderActionAlignment{
		"center": AddContentHeaderActionAlignment_Center,
		"left":   AddContentHeaderActionAlignment_Left,
		"right":  AddContentHeaderActionAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddContentHeaderActionAlignment(input)
	return &out, nil
}
