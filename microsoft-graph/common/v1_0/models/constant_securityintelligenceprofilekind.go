package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileKind string

const (
	SecurityIntelligenceProfileKindactor SecurityIntelligenceProfileKind = "Actor"
	SecurityIntelligenceProfileKindtool  SecurityIntelligenceProfileKind = "Tool"
)

func PossibleValuesForSecurityIntelligenceProfileKind() []string {
	return []string{
		string(SecurityIntelligenceProfileKindactor),
		string(SecurityIntelligenceProfileKindtool),
	}
}

func parseSecurityIntelligenceProfileKind(input string) (*SecurityIntelligenceProfileKind, error) {
	vals := map[string]SecurityIntelligenceProfileKind{
		"actor": SecurityIntelligenceProfileKindactor,
		"tool":  SecurityIntelligenceProfileKindtool,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIntelligenceProfileKind(input)
	return &out, nil
}
