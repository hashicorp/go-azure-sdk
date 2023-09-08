package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFormattedContentFormat string

const (
	SecurityFormattedContentFormathtml     SecurityFormattedContentFormat = "Html"
	SecurityFormattedContentFormatmarkdown SecurityFormattedContentFormat = "Markdown"
	SecurityFormattedContentFormattext     SecurityFormattedContentFormat = "Text"
)

func PossibleValuesForSecurityFormattedContentFormat() []string {
	return []string{
		string(SecurityFormattedContentFormathtml),
		string(SecurityFormattedContentFormatmarkdown),
		string(SecurityFormattedContentFormattext),
	}
}

func parseSecurityFormattedContentFormat(input string) (*SecurityFormattedContentFormat, error) {
	vals := map[string]SecurityFormattedContentFormat{
		"html":     SecurityFormattedContentFormathtml,
		"markdown": SecurityFormattedContentFormatmarkdown,
		"text":     SecurityFormattedContentFormattext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFormattedContentFormat(input)
	return &out, nil
}
