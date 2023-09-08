package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalItemContentType string

const (
	ExternalItemContentTypehtml ExternalItemContentType = "Html"
	ExternalItemContentTypetext ExternalItemContentType = "Text"
)

func PossibleValuesForExternalItemContentType() []string {
	return []string{
		string(ExternalItemContentTypehtml),
		string(ExternalItemContentTypetext),
	}
}

func parseExternalItemContentType(input string) (*ExternalItemContentType, error) {
	vals := map[string]ExternalItemContentType{
		"html": ExternalItemContentTypehtml,
		"text": ExternalItemContentTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalItemContentType(input)
	return &out, nil
}
