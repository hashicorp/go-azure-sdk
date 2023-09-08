package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventQueryQueryType string

const (
	SecurityEventQueryQueryTypefiles    SecurityEventQueryQueryType = "Files"
	SecurityEventQueryQueryTypemessages SecurityEventQueryQueryType = "Messages"
)

func PossibleValuesForSecurityEventQueryQueryType() []string {
	return []string{
		string(SecurityEventQueryQueryTypefiles),
		string(SecurityEventQueryQueryTypemessages),
	}
}

func parseSecurityEventQueryQueryType(input string) (*SecurityEventQueryQueryType, error) {
	vals := map[string]SecurityEventQueryQueryType{
		"files":    SecurityEventQueryQueryTypefiles,
		"messages": SecurityEventQueryQueryTypemessages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventQueryQueryType(input)
	return &out, nil
}
