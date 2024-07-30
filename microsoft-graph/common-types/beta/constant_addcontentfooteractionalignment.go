package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddContentFooterActionAlignment string

const (
	AddContentFooterActionAlignment_Center AddContentFooterActionAlignment = "center"
	AddContentFooterActionAlignment_Left   AddContentFooterActionAlignment = "left"
	AddContentFooterActionAlignment_Right  AddContentFooterActionAlignment = "right"
)

func PossibleValuesForAddContentFooterActionAlignment() []string {
	return []string{
		string(AddContentFooterActionAlignment_Center),
		string(AddContentFooterActionAlignment_Left),
		string(AddContentFooterActionAlignment_Right),
	}
}

func (s *AddContentFooterActionAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAddContentFooterActionAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAddContentFooterActionAlignment(input string) (*AddContentFooterActionAlignment, error) {
	vals := map[string]AddContentFooterActionAlignment{
		"center": AddContentFooterActionAlignment_Center,
		"left":   AddContentFooterActionAlignment_Left,
		"right":  AddContentFooterActionAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddContentFooterActionAlignment(input)
	return &out, nil
}
