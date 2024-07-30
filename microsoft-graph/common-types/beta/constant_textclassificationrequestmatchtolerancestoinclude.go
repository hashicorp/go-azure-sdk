package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequestMatchTolerancesToInclude string

const (
	TextClassificationRequestMatchTolerancesToInclude_Exact TextClassificationRequestMatchTolerancesToInclude = "exact"
	TextClassificationRequestMatchTolerancesToInclude_Near  TextClassificationRequestMatchTolerancesToInclude = "near"
)

func PossibleValuesForTextClassificationRequestMatchTolerancesToInclude() []string {
	return []string{
		string(TextClassificationRequestMatchTolerancesToInclude_Exact),
		string(TextClassificationRequestMatchTolerancesToInclude_Near),
	}
}

func (s *TextClassificationRequestMatchTolerancesToInclude) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTextClassificationRequestMatchTolerancesToInclude(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTextClassificationRequestMatchTolerancesToInclude(input string) (*TextClassificationRequestMatchTolerancesToInclude, error) {
	vals := map[string]TextClassificationRequestMatchTolerancesToInclude{
		"exact": TextClassificationRequestMatchTolerancesToInclude_Exact,
		"near":  TextClassificationRequestMatchTolerancesToInclude_Near,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TextClassificationRequestMatchTolerancesToInclude(input)
	return &out, nil
}
