package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemBodyContentType string

const (
	ItemBodyContentType_Html ItemBodyContentType = "html"
	ItemBodyContentType_Text ItemBodyContentType = "text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentType_Html),
		string(ItemBodyContentType_Text),
	}
}

func (s *ItemBodyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemBodyContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentType_Html,
		"text": ItemBodyContentType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}
