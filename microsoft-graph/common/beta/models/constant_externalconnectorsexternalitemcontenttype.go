package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalItemContentType string

const (
	ExternalConnectorsExternalItemContentTypehtml ExternalConnectorsExternalItemContentType = "Html"
	ExternalConnectorsExternalItemContentTypetext ExternalConnectorsExternalItemContentType = "Text"
)

func PossibleValuesForExternalConnectorsExternalItemContentType() []string {
	return []string{
		string(ExternalConnectorsExternalItemContentTypehtml),
		string(ExternalConnectorsExternalItemContentTypetext),
	}
}

func parseExternalConnectorsExternalItemContentType(input string) (*ExternalConnectorsExternalItemContentType, error) {
	vals := map[string]ExternalConnectorsExternalItemContentType{
		"html": ExternalConnectorsExternalItemContentTypehtml,
		"text": ExternalConnectorsExternalItemContentTypetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsExternalItemContentType(input)
	return &out, nil
}
