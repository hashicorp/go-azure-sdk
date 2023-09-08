package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileIndicatorSource string

const (
	SecurityIntelligenceProfileIndicatorSourcemicrosoft SecurityIntelligenceProfileIndicatorSource = "Microsoft"
	SecurityIntelligenceProfileIndicatorSourceosint     SecurityIntelligenceProfileIndicatorSource = "Osint"
	SecurityIntelligenceProfileIndicatorSourcepublic    SecurityIntelligenceProfileIndicatorSource = "Public"
)

func PossibleValuesForSecurityIntelligenceProfileIndicatorSource() []string {
	return []string{
		string(SecurityIntelligenceProfileIndicatorSourcemicrosoft),
		string(SecurityIntelligenceProfileIndicatorSourceosint),
		string(SecurityIntelligenceProfileIndicatorSourcepublic),
	}
}

func parseSecurityIntelligenceProfileIndicatorSource(input string) (*SecurityIntelligenceProfileIndicatorSource, error) {
	vals := map[string]SecurityIntelligenceProfileIndicatorSource{
		"microsoft": SecurityIntelligenceProfileIndicatorSourcemicrosoft,
		"osint":     SecurityIntelligenceProfileIndicatorSourceosint,
		"public":    SecurityIntelligenceProfileIndicatorSourcepublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIntelligenceProfileIndicatorSource(input)
	return &out, nil
}
