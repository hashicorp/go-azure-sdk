package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemEmailType string

const (
	ItemEmailType_Main     ItemEmailType = "main"
	ItemEmailType_Other    ItemEmailType = "other"
	ItemEmailType_Personal ItemEmailType = "personal"
	ItemEmailType_Unknown  ItemEmailType = "unknown"
	ItemEmailType_Work     ItemEmailType = "work"
)

func PossibleValuesForItemEmailType() []string {
	return []string{
		string(ItemEmailType_Main),
		string(ItemEmailType_Other),
		string(ItemEmailType_Personal),
		string(ItemEmailType_Unknown),
		string(ItemEmailType_Work),
	}
}

func (s *ItemEmailType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemEmailType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemEmailType(input string) (*ItemEmailType, error) {
	vals := map[string]ItemEmailType{
		"main":     ItemEmailType_Main,
		"other":    ItemEmailType_Other,
		"personal": ItemEmailType_Personal,
		"unknown":  ItemEmailType_Unknown,
		"work":     ItemEmailType_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailType(input)
	return &out, nil
}
