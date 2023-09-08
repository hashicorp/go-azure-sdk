package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityArticleIndicatorSource string

const (
	SecurityArticleIndicatorSourcemicrosoft SecurityArticleIndicatorSource = "Microsoft"
	SecurityArticleIndicatorSourceosint     SecurityArticleIndicatorSource = "Osint"
	SecurityArticleIndicatorSourcepublic    SecurityArticleIndicatorSource = "Public"
)

func PossibleValuesForSecurityArticleIndicatorSource() []string {
	return []string{
		string(SecurityArticleIndicatorSourcemicrosoft),
		string(SecurityArticleIndicatorSourceosint),
		string(SecurityArticleIndicatorSourcepublic),
	}
}

func parseSecurityArticleIndicatorSource(input string) (*SecurityArticleIndicatorSource, error) {
	vals := map[string]SecurityArticleIndicatorSource{
		"microsoft": SecurityArticleIndicatorSourcemicrosoft,
		"osint":     SecurityArticleIndicatorSourceosint,
		"public":    SecurityArticleIndicatorSourcepublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityArticleIndicatorSource(input)
	return &out, nil
}
