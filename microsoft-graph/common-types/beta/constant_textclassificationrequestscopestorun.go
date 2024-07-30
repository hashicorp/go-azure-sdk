package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequestScopesToRun string

const (
	TextClassificationRequestScopesToRun_FullDocument    TextClassificationRequestScopesToRun = "fullDocument"
	TextClassificationRequestScopesToRun_PartialDocument TextClassificationRequestScopesToRun = "partialDocument"
)

func PossibleValuesForTextClassificationRequestScopesToRun() []string {
	return []string{
		string(TextClassificationRequestScopesToRun_FullDocument),
		string(TextClassificationRequestScopesToRun_PartialDocument),
	}
}

func (s *TextClassificationRequestScopesToRun) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTextClassificationRequestScopesToRun(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTextClassificationRequestScopesToRun(input string) (*TextClassificationRequestScopesToRun, error) {
	vals := map[string]TextClassificationRequestScopesToRun{
		"fulldocument":    TextClassificationRequestScopesToRun_FullDocument,
		"partialdocument": TextClassificationRequestScopesToRun_PartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TextClassificationRequestScopesToRun(input)
	return &out, nil
}
