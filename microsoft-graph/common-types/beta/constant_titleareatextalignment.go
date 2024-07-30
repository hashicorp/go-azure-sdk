package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaTextAlignment string

const (
	TitleAreaTextAlignment_Center TitleAreaTextAlignment = "center"
	TitleAreaTextAlignment_Left   TitleAreaTextAlignment = "left"
)

func PossibleValuesForTitleAreaTextAlignment() []string {
	return []string{
		string(TitleAreaTextAlignment_Center),
		string(TitleAreaTextAlignment_Left),
	}
}

func (s *TitleAreaTextAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTitleAreaTextAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTitleAreaTextAlignment(input string) (*TitleAreaTextAlignment, error) {
	vals := map[string]TitleAreaTextAlignment{
		"center": TitleAreaTextAlignment_Center,
		"left":   TitleAreaTextAlignment_Left,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaTextAlignment(input)
	return &out, nil
}
