package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormattedContentFormat string

const (
	FormattedContentFormat_Html     FormattedContentFormat = "html"
	FormattedContentFormat_Markdown FormattedContentFormat = "markdown"
	FormattedContentFormat_Text     FormattedContentFormat = "text"
)

func PossibleValuesForFormattedContentFormat() []string {
	return []string{
		string(FormattedContentFormat_Html),
		string(FormattedContentFormat_Markdown),
		string(FormattedContentFormat_Text),
	}
}

func (s *FormattedContentFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFormattedContentFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFormattedContentFormat(input string) (*FormattedContentFormat, error) {
	vals := map[string]FormattedContentFormat{
		"html":     FormattedContentFormat_Html,
		"markdown": FormattedContentFormat_Markdown,
		"text":     FormattedContentFormat_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FormattedContentFormat(input)
	return &out, nil
}
