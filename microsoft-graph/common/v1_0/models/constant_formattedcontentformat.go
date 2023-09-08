package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormattedContentFormat string

const (
	FormattedContentFormathtml     FormattedContentFormat = "Html"
	FormattedContentFormatmarkdown FormattedContentFormat = "Markdown"
	FormattedContentFormattext     FormattedContentFormat = "Text"
)

func PossibleValuesForFormattedContentFormat() []string {
	return []string{
		string(FormattedContentFormathtml),
		string(FormattedContentFormatmarkdown),
		string(FormattedContentFormattext),
	}
}

func parseFormattedContentFormat(input string) (*FormattedContentFormat, error) {
	vals := map[string]FormattedContentFormat{
		"html":     FormattedContentFormathtml,
		"markdown": FormattedContentFormatmarkdown,
		"text":     FormattedContentFormattext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FormattedContentFormat(input)
	return &out, nil
}
