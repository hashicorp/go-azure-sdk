package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemBodyContentType string

const (
	ItemBodyContentTypehtml ItemBodyContentType = "Html"
	ItemBodyContentTypetext ItemBodyContentType = "Text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentTypehtml),
		string(ItemBodyContentTypetext),
	}
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentTypehtml,
		"text": ItemBodyContentTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}
