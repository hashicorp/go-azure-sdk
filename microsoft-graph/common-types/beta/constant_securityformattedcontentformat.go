package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFormattedContentFormat string

const (
	SecurityFormattedContentFormat_Html     SecurityFormattedContentFormat = "html"
	SecurityFormattedContentFormat_Markdown SecurityFormattedContentFormat = "markdown"
	SecurityFormattedContentFormat_Text     SecurityFormattedContentFormat = "text"
)

func PossibleValuesForSecurityFormattedContentFormat() []string {
	return []string{
		string(SecurityFormattedContentFormat_Html),
		string(SecurityFormattedContentFormat_Markdown),
		string(SecurityFormattedContentFormat_Text),
	}
}

func (s *SecurityFormattedContentFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFormattedContentFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFormattedContentFormat(input string) (*SecurityFormattedContentFormat, error) {
	vals := map[string]SecurityFormattedContentFormat{
		"html":     SecurityFormattedContentFormat_Html,
		"markdown": SecurityFormattedContentFormat_Markdown,
		"text":     SecurityFormattedContentFormat_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFormattedContentFormat(input)
	return &out, nil
}
