package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryFileSourceType string

const (
	SecurityEdiscoveryFileSourceTypemailbox SecurityEdiscoveryFileSourceType = "Mailbox"
	SecurityEdiscoveryFileSourceTypesite    SecurityEdiscoveryFileSourceType = "Site"
)

func PossibleValuesForSecurityEdiscoveryFileSourceType() []string {
	return []string{
		string(SecurityEdiscoveryFileSourceTypemailbox),
		string(SecurityEdiscoveryFileSourceTypesite),
	}
}

func parseSecurityEdiscoveryFileSourceType(input string) (*SecurityEdiscoveryFileSourceType, error) {
	vals := map[string]SecurityEdiscoveryFileSourceType{
		"mailbox": SecurityEdiscoveryFileSourceTypemailbox,
		"site":    SecurityEdiscoveryFileSourceTypesite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoveryFileSourceType(input)
	return &out, nil
}
