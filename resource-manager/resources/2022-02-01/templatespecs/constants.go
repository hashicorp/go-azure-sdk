package templatespecs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateSpecExpandKind string

const (
	TemplateSpecExpandKindVersions TemplateSpecExpandKind = "versions"
)

func PossibleValuesForTemplateSpecExpandKind() []string {
	return []string{
		string(TemplateSpecExpandKindVersions),
	}
}

func (s *TemplateSpecExpandKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTemplateSpecExpandKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTemplateSpecExpandKind(input string) (*TemplateSpecExpandKind, error) {
	vals := map[string]TemplateSpecExpandKind{
		"versions": TemplateSpecExpandKindVersions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemplateSpecExpandKind(input)
	return &out, nil
}
